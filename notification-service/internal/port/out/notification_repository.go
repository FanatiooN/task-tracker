package out

import (
	"context"
	"task-tracker/notification-service/internal/domain"

	"github.com/google/uuid"
)

type NotificationRepository interface {
	Save(ctx context.Context, notification domain.Notification) (uuid.UUID, error)
	GetStatistics(ctx context.Context, filter domain.StatisticsFilter) ([]domain.Statistics, error)
}
