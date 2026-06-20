package out

import (
	"context"
	"task-tracker/auth-service/internal/domain"

	"github.com/google/uuid"
)

type TokenRepository interface {
	Save(ctx context.Context, token domain.RefreshToken) error
	FindByUserID(ctx context.Context, userID uuid.UUID) (domain.RefreshToken, error)
}
