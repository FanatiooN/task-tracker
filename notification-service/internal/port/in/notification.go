package in

import (
	"context"
	"task-tracker/notification-service/internal/domain"
)

type NotificationService interface {
	Send(ctx context.Context, notification domain.Notification) error
	GetStats(ctx context.Context, filter domain.StatisticsFilter) ([]domain.Statistics, error)
}
