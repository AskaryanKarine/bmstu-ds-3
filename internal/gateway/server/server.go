package server

import (
	"context"
	"fmt"
	"github.com/AskaryanKarine/bmstu-ds-3/internal/gateway/clients"
	"github.com/AskaryanKarine/bmstu-ds-3/internal/gateway/config"
	"github.com/AskaryanKarine/bmstu-ds-3/internal/gateway/retryer"
	"github.com/AskaryanKarine/bmstu-ds-3/pkg/app"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	circuit "github.com/rubyist/circuitbreaker"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	gracefulShutdownDeadline = 10 * time.Second
)

type Server struct {
	echo *echo.Echo
	cfg  config.Config

	loyalty     loyaltyClient
	payment     paymentClient
	reservation reservationClient

	retryerQueue retryerQueueProducer
}

func NewServer(cfg config.Config, rtq retryerQueueProducer) *Server {
	e := echo.New()
	payment := clients.NewPaymentClient(circuit.NewHTTPClient(0, 10, nil), cfg.PaymentService)
	reservation := clients.NewReservationClient(circuit.NewHTTPClient(0, 10, nil), cfg.ReservationService)
	loyalty := clients.NewLoyaltyClient(circuit.NewHTTPClient(0, 10, nil), cfg.LoyaltyService)
	s := &Server{
		echo:         e,
		cfg:          cfg,
		loyalty:      loyalty,
		payment:      payment,
		reservation:  reservation,
		retryerQueue: rtq,
	}

	app.SetStandardSetting(e)
	app.AddHealthCheck(e)

	api := s.echo.Group("/api/v1")

	api.GET("/hotels", s.getHotels)

	api.GET("/me", s.getUserInfo, app.GetUsernameMW())

	reservations := api.Group("/reservations", app.GetUsernameMW())
	reservations.GET("", s.getReservations, app.GetUsernameMW())
	reservations.POST("", s.createReservation, app.GetUsernameMW())
	reservations.GET("/:uid", s.getReservationsByUID, app.GetUsernameMW())
	reservations.DELETE("/:uid", s.canceledReservation, app.GetUsernameMW())

	api.GET("/loyalty", s.getLoyalty, app.GetUsernameMW())

	return s
}

func (s *Server) Run(port int) {
	portStr := fmt.Sprintf(":%d", port)
	go func() {
		log.Info("server starting on", "port", portStr)
		if err := s.echo.Start(portStr); err != nil {
			log.Fatal(err)
		}
	}()

	consumer := retryer.NewLoyaltyRetryerConsumer(s.cfg, s.loyalty)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Info("Loyalty Consumer started")
	go consumer.RetryLoyaltyDecrease(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	cancel()
	if err := consumer.Stop(); err != nil {
		log.Error(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), gracefulShutdownDeadline)
	defer cancel()

	log.Info("server shutting down")
	if err := s.echo.Server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
