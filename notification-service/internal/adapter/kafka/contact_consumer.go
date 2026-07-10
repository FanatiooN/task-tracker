package kafka

import (
	"context"
	"encoding/json"
	kafkaevents "task-tracker/kafka"
	"task-tracker/notification-service/internal/port/out"
	pkgkafka "task-tracker/pkg/kafka"
)

type ContactConsumer struct {
	consumer        *pkgkafka.Consumer
	userContactRepo out.UserContactRepository
}

func NewContactConsumer(consumer *pkgkafka.Consumer, userContactRepo out.UserContactRepository) *ContactConsumer {
	return &ContactConsumer{
		consumer:        consumer,
		userContactRepo: userContactRepo,
	}
}

func (c *ContactConsumer) Start(ctx context.Context) {
	c.consumer.Consume(ctx, func(msg []byte) error {
		var event kafkaevents.UserContactLinkedEvent
		if err := json.Unmarshal(msg, &event); err != nil {
			return err
		}

		return c.userContactRepo.Save(ctx, event.UserID, event.Provider, event.Contact)
	})
}
