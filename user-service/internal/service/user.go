package service

import (
	"context"
	"errors"
	"strings"
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
	user.Name = strings.TrimSpace(user.Name)
	if user.Name == "" {
		return domain.User{}, errors.New("invalid name")
	}

	var email string

	if user.Email != nil {
		email = strings.TrimSpace(*user.Email)
		if !validEmail.MatchString(email) {
			return domain.User{}, errors.New("invalid email")
		}

		user.Email = &email
	}

	response, err := u.repository.Save(ctx, domain.User{
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return domain.User{}, err
	}

	return response, nil
}

func (u UserService) GetUser(ctx context.Context, id uuid.UUID) (domain.User, error) {
	response, err := u.repository.FindByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return response, nil
}

func (u UserService) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	response, err := u.repository.FindByID(ctx, user.ID)
	if err != nil {
		return domain.User{}, err
	}

	user.Name = strings.TrimSpace(user.Name)
	if user.Name != "" {
		response.Name = user.Name
	}

	var email string

	if user.Email != nil {
		email = strings.TrimSpace(*user.Email)
		if !validEmail.MatchString(email) {
			return domain.User{}, errors.New("invalid email")
		}

		response.Email = &email
	}

	response, err = u.repository.Update(ctx, response)
	if err != nil {
		return domain.User{}, err
	}

	return response, nil
}

func (u UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
