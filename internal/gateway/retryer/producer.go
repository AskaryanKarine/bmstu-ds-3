package retryer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AskaryanKarine/bmstu-ds-3/internal/gateway/config"
	"github.com/AskaryanKarine/bmstu-ds-3/internal/gateway/models"
	"github.com/segmentio/kafka-go"
	"net"
	"strconv"
)

type Writer interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) error
}

type QueueRetryerProducer struct {
	writer Writer

	host                 string
	decreaseLoyaltyTopic string
}

func NewQueueRetryerProducer(cfg config.Config) (*QueueRetryerProducer, error) {
	writer := kafka.Writer{
		Addr: kafka.TCP(cfg.KafkaHost),
	}
	p := QueueRetryerProducer{
		host:                 cfg.KafkaHost,
		decreaseLoyaltyTopic: cfg.KafkaLoyaltyDecreaseTopic,
		writer:               &writer,
	}

	err := p.createMissedTopic()
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (p *QueueRetryerProducer) createMissedTopic() error {
	conn, err := kafka.Dial("tcp", p.host)
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}

	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             p.decreaseLoyaltyTopic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		return err
	}

	return nil
}

func (p *QueueRetryerProducer) RetryLoyaltyDecrease(ctx context.Context, username string) error {
	msg := models.DecreaseLoyaltyMsg{Username: username}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("could not marshal decrease loyalty msg: %w", err)
	}

	err = p.writer.WriteMessages(ctx, kafka.Message{
		Topic: p.decreaseLoyaltyTopic,
		Value: msgBytes,
	})
	if err != nil {
		return fmt.Errorf("could not write messages: %w", err)
	}
	return nil
}
