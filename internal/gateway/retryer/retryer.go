package retryer

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/AskaryanKarine/bmstu-ds-3/internal/gateway/config"
	"github.com/AskaryanKarine/bmstu-ds-3/internal/gateway/models"
	"github.com/labstack/gommon/log"
	"github.com/segmentio/kafka-go"
	"time"
)

type LoyaltyRetryerConsumer struct {
	topic string

	client   loyaltyClient
	consumer Reader
}

type Reader interface {
	FetchMessage(ctx context.Context) (kafka.Message, error)
	CommitMessages(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

type loyaltyClient interface {
	DecreaseLoyalty(username string) error
}

func NewLoyaltyRetryerConsumer(cfg config.Config, client loyaltyClient) *LoyaltyRetryerConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.KafkaHost},
		Topic:   cfg.KafkaLoyaltyDecreaseTopic,
		GroupID: cfg.KafkaLoyaltyDecreaseTopic,
	})

	return &LoyaltyRetryerConsumer{
		topic:    cfg.KafkaLoyaltyDecreaseTopic,
		consumer: reader,
		client:   client,
	}
}

func (r *LoyaltyRetryerConsumer) Stop() error {
	return r.consumer.Close()
}

func (r *LoyaltyRetryerConsumer) RetryLoyaltyDecrease(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Info("Consumer shutdown")
			return
		default:
			msg, err := r.consumer.FetchMessage(ctx)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					log.Info("Consumer context canceled, exiting")
					return
				}
				log.Error("can not fetch message: ", err)
				continue
			}
			log.Info("Received message: ", string(msg.Value))
			var usernameMsg models.DecreaseLoyaltyMsg
			err = json.Unmarshal(msg.Value, &usernameMsg)
			if err != nil {
				log.Error("can not unmarshal message ", err)
				continue
			}
			for {
				err = r.client.DecreaseLoyalty(usernameMsg.Username)
				if err == nil {
					log.Info("Successful request")
					break
				}
				log.Info("Loyalty Service until unavailable, sleep 10 sec")
				time.Sleep(10 * time.Second)
			}

			err = r.consumer.CommitMessages(ctx, msg)
			if err != nil {
				log.Error("can not commit message ", err)
				continue
			}
		}
	}
}
