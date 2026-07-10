package out

import (
	"context"

	"github.com/google/uuid"
)

type ContactProducer interface {
	Produce(ctx context.Context, userID uuid.UUID, provider, contact string) error
}
