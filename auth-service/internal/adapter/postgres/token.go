package postgres

import (
	"context"
	"task-tracker/auth-service/internal/db"
	"task-tracker/auth-service/internal/domain"

	"github.com/google/uuid"
)

type TokenRepository struct {
	queries *db.Queries
}

func NewTokenRepository(queries *db.Queries) *TokenRepository {
	return &TokenRepository{queries: queries}
}

func (t TokenRepository) Save(ctx context.Context, token domain.RefreshToken) error {
	params := db.CreateTokenParams{
		UserID:    token.UserID,
		TokenHash: token.TokenHash,
		ExpiresAt: token.ExpiresAt,
	}

	_, err := t.queries.CreateToken(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (t TokenRepository) FindByUserID(ctx context.Context, userID uuid.UUID) (domain.RefreshToken, error) {
	row, err := t.queries.FindByUserID(ctx, userID)
	if err != nil {
		return domain.RefreshToken{}, err
	}

	return domain.RefreshToken{
		ID:        row.ID,
		UserID:    row.UserID,
		TokenHash: row.TokenHash,
		ExpiresAt: row.ExpiresAt,
		IsRevoked: row.IsRevoked.Bool,
	}, nil
}

func (t TokenRepository) FindByTokenHash(ctx context.Context, tokenHash string) (domain.RefreshToken, error) {
	row, err := t.queries.FindByTokenHash(ctx, tokenHash)
	if err != nil {
		return domain.RefreshToken{}, err
	}

	return domain.RefreshToken{
		ID:        row.ID,
		UserID:    row.UserID,
		TokenHash: row.TokenHash,
		ExpiresAt: row.ExpiresAt,
		IsRevoked: row.IsRevoked.Bool,
	}, nil
}

func (t TokenRepository) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	err := t.queries.DeleteByUserID(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
