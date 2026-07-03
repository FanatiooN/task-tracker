package out

import (
	"context"
	"task-tracker/notification-service/internal/domain"
)

type NotificationProvider interface {
	Send(context context.Context, notification domain.Notification) error
}
