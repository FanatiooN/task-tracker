package domain

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID               uuid.UUID
	UserID           uuid.UUID
	Type             string
	Provider         string
	NotificationBody NotificationBody
	CreatedAt        time.Time
}

type NotificationBody struct {
	Text      string
	PhotoUrls []string
}

type Statistics struct {
	Type       string
	Provider   string
	Date       time.Time
	UniqueHits int
	TotalHits  int
}

type StatisticsFilter struct {
	Date     *time.Time
	Provider *string
}
