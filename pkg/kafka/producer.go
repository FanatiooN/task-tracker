package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(topic, broker string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:  kafka.TCP(broker),
			Topic: topic,
		},
	}
}

func (p Producer) Produce(ctx context.Context, key, value []byte) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
	}

	err := p.writer.WriteMessages(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}

func (p Producer) Close() error {
	err := p.writer.Close()
	return err
}
