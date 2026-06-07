package postgres

import (
	"context"
	"task-tracker/user-service/internal/db"
	"task-tracker/user-service/internal/domain"

	"github.com/google/uuid"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(queries *db.Queries) *UserRepository {
	return &UserRepository{queries: queries}
}

func (u UserRepository) Save(ctx context.Context, user domain.User) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) FindByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Update(ctx context.Context, user domain.User) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
