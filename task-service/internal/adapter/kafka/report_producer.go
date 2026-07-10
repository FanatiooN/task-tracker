package kafka

import (
	"context"
	"encoding/json"
	"log"
	kafkaevents "task-tracker/kafka"
	pkgkafka "task-tracker/pkg/kafka"

	"github.com/google/uuid"
)

type ReportProducer struct {
	producer *pkgkafka.Producer
}

func NewReportProducer(producer *pkgkafka.Producer) *ReportProducer {
	return &ReportProducer{
		producer: producer,
	}
}

func (p *ReportProducer) Produce(ctx context.Context, userID uuid.UUID, notificationType, provider, text string, photoUrls []string) error {
	event := kafkaevents.SendNotificationEvent{
		UserID:   userID,
		Type:     notificationType,
		Provider: provider,
		NotificationBody: struct {
			Text      string   `json:"text"`
			PhotoUrls []string `json:"photo_urls"`
		}{
			Text:      text,
			PhotoUrls: photoUrls,
		},
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
