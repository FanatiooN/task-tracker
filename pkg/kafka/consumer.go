package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(topic, groupID string, brokers ...string) *Consumer {
	conf := kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	}

	return &Consumer{
		reader: kafka.NewReader(conf),
	}
}

func (c Consumer) Consume(ctx context.Context, handler func([]byte) error) {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			return
		}

		err = handler(msg.Value)
		if err != nil {
			log.Printf("handler error: %v", err)
			continue
		}
	}
}

func (c Consumer) Close() error {
	err := c.reader.Close()
	return err
}
