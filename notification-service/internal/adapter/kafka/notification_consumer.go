package kafka

import (
	"context"
	"encoding/json"
	kafkaevents "task-tracker/kafka"
	"task-tracker/notification-service/internal/domain"
	"task-tracker/notification-service/internal/service"
	pkgkafka "task-tracker/pkg/kafka"
)

type NotificationConsumer struct {
	consumer            *pkgkafka.Consumer
	notificationService *service.NotificationService
}

func NewNotificationConsumer(consumer *pkgkafka.Consumer, notificationService *service.NotificationService) *NotificationConsumer {
	return &NotificationConsumer{
		consumer:            consumer,
		notificationService: notificationService,
	}
}

func (c *NotificationConsumer) Start(ctx context.Context) {
	c.consumer.Consume(ctx, func(msg []byte) error {
		var event kafkaevents.SendNotificationEvent
		if err := json.Unmarshal(msg, &event); err != nil {
			return err
		}

		notification := domain.Notification{
			UserID:   event.UserID,
			Type:     event.Type,
			Provider: event.Provider,
			NotificationBody: domain.NotificationBody{
				Text:      event.NotificationBody.Text,
				PhotoUrls: event.NotificationBody.PhotoUrls,
			},
		}

		return c.notificationService.Send(ctx, notification)
	})
}
