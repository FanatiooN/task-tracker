package kafka

import (
	"context"
	"encoding/json"
	"log"
	kafkaevents "task-tracker/kafka"
	pkgkafka "task-tracker/pkg/kafka"

	"github.com/google/uuid"
)

type ContactProducer struct {
	producer *pkgkafka.Producer
}

func NewContactProducer(producer *pkgkafka.Producer) *ContactProducer {
	return &ContactProducer{
		producer: producer,
	}
}

func (p *ContactProducer) Produce(ctx context.Context, userID uuid.UUID, provider, contact string) error {
	event := kafkaevents.UserContactLinkedEvent{
		UserID:   userID,
		Provider: provider,
		Contact:  contact,
	}

	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	key := userID.String()

	err = p.producer.Produce(ctx, []byte(key), eventJSON)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	return nil
}
