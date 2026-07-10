package kafka

import (
	"github.com/google/uuid"
)

type UserContactLinkedEvent struct {
	UserID   uuid.UUID `json:"user_id"`
	Provider string    `json:"provider"`
	Contact  string    `json:"contact"`
}

type SendNotificationEvent struct {
	UserID           uuid.UUID `json:"user_id"`
	Type             string    `json:"type"`
	Provider         string    `json:"provider"`
	NotificationBody struct {
		Text      string   `json:"text"`
		PhotoUrls []string `json:"photo_urls"`
	} `json:"notification_body"`
}
