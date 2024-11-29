package main

import (
	"github.com/AskaryanKarine/bmstu-ds-3/internal/gateway/config"
	"github.com/AskaryanKarine/bmstu-ds-3/internal/gateway/retryer"
	"github.com/AskaryanKarine/bmstu-ds-3/internal/gateway/server"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	rqp, err := retryer.NewQueueRetryerProducer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	s := server.NewServer(cfg, rqp)
	s.Run(cfg.Port)
}
