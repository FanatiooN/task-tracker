package in

import (
	"context"
	"task-tracker/user-service/internal/domain"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, id uuid.UUID) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) (domain.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
