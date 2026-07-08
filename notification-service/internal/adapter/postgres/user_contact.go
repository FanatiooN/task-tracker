package postgres

import (
	"context"
	"task-tracker/notification-service/internal/db"

	"github.com/google/uuid"
)

type UserContactRepository struct {
	queries *db.Queries
}

func NewUserContactRepository(queries *db.Queries) *UserContactRepository {
	return &UserContactRepository{queries: queries}
}

func (u UserContactRepository) Save(ctx context.Context, userID uuid.UUID, provider, contact string) error {
	params := db.SaveContactParams{
		UserID:   userID,
		Provider: provider,
		Contact:  contact,
	}

	_, err := u.queries.SaveContact(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (u UserContactRepository) GetContact(ctx context.Context, userID uuid.UUID, provider string) (string, error) {
	params := db.GetContactParams{
		UserID:   userID,
		Provider: provider,
	}

	contact, err := u.queries.GetContact(ctx, params)
	if err != nil {
		return "", err
	}

	return contact, nil
}
