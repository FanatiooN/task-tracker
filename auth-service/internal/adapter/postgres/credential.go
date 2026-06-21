package postgres

import (
	"context"
	"task-tracker/auth-service/internal/db"
	"task-tracker/auth-service/internal/domain"
)

type CredentialRepository struct {
	queries *db.Queries
}

func NewCredentialRepository(queries *db.Queries) *CredentialRepository {
	return &CredentialRepository{queries: queries}
}

func (c CredentialRepository) Save(ctx context.Context, credential domain.Credential) error {
	params := db.CreateCredentialsParams{
		UserID:       credential.UserID,
		Email:        credential.Email,
		PasswordHash: credential.PasswordHash,
	}

	_, err := c.queries.CreateCredentials(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (c CredentialRepository) FindByEmail(ctx context.Context, email string) (domain.Credential, error) {
	row, err := c.queries.FindByEmail(ctx, email)
	if err != nil {
		return domain.Credential{}, err
	}

	return domain.Credential{
		UserID:       row.UserID,
		Email:        row.Email,
		PasswordHash: row.PasswordHash,
	}, nil
}
