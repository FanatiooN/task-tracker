package postgres

import (
	"context"
	"task-tracker/user-service/internal/db"
	"task-tracker/user-service/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(queries *db.Queries) *UserRepository {
	return &UserRepository{queries: queries}
}

func (u *UserRepository) Save(ctx context.Context, user domain.User) (domain.User, error) {
	var params db.CreateUserParams

	params.Name = user.Name

	if user.Email != nil {
		params.Email = pgtype.Text{
			String: *user.Email,
			Valid:  true,
		}
	} else {
		params.Email = pgtype.Text{
			String: "",
			Valid:  false,
		}
	}

	row, err := u.queries.CreateUser(ctx, params)
	if err != nil {
		return domain.User{}, err
	}

	var email *string
	if row.Email.Valid {
		email = &row.Email.String
	}

	return domain.User{
		ID:    row.ID,
		Name:  row.Name,
		Email: email,
	}, nil
}

func (u UserRepository) FindByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	row, err := u.queries.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	var email *string
	if row.Email.Valid {
		email = &row.Email.String
	}

	return domain.User{
		ID:    row.ID,
		Name:  row.Name,
		Email: email,
	}, nil
}

func (u UserRepository) Update(ctx context.Context, user domain.User) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
