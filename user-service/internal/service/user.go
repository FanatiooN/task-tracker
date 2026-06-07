package service

import (
	"context"
	"task-tracker/user-service/internal/domain"
	"task-tracker/user-service/internal/port/out"

	"github.com/google/uuid"
)

type UserService struct {
	repository out.UserRepository
}

func NewUserService(repository out.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (u UserService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetUser(ctx context.Context, id uuid.UUID) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
