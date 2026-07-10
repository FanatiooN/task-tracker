package out

import (
	"context"

	"github.com/google/uuid"
)

type ReportProducer interface {
	Produce(ctx context.Context, userID uuid.UUID, notificationType, provider, text string, photoUrls []string) error
}
