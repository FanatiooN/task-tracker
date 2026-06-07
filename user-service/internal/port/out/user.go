package out

import (
	"context"
	"task-tracker/user-service/internal/domain"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) (domain.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (domain.User, error)
	Update(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
