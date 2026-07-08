package out

import (
	"context"

	"github.com/google/uuid"
)

type UserContactRepository interface {
	Save(ctx context.Context, userID uuid.UUID, provider, contact string) error
	GetContact(ctx context.Context, userID uuid.UUID, provider string) (string, error)
}
