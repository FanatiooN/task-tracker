package kafka

import (
	"context"
	"net"
	"strconv"

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

func (p *Producer) Produce(ctx context.Context, key, value []byte) error {
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

func (p *Producer) Close() error {
	err := p.writer.Close()
	return err
}

func CreateTopic(topic, broker string) error {
	conn, err := kafka.Dial("tcp", broker)
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return err
	}
	defer controllerConn.Close()

	err = controllerConn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})

	return err
}
