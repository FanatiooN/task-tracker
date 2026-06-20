package out

import (
	"context"
	"task-tracker/auth-service/internal/domain"

	"github.com/google/uuid"
)

type TokenRepository interface {
	Save(ctx context.Context, token domain.RefreshToken)
	FindByUserID(ctx context.Context, userID uuid.UUID)
}
