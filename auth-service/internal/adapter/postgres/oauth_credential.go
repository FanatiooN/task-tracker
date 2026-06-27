package postgres

import (
	"context"
	"task-tracker/auth-service/internal/db"
	"task-tracker/auth-service/internal/domain"
)

type OAuthCredential struct {
	queries *db.Queries
}

func NewOAuthCredential(queries *db.Queries) *OAuthCredential {
	return &OAuthCredential{queries: queries}
}

func (O OAuthCredential) Save(ctx context.Context, credential domain.OAuthCredential) error {
	params := db.CreateOAuthCredentialParams{
		UserID:     credential.UserID,
		Provider:   credential.Provider,
		ProviderID: credential.ProviderID,
	}

	_, err := O.queries.CreateOAuthCredential(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (O OAuthCredential) FindByProvider(ctx context.Context, provider, providerID string) (domain.OAuthCredential, error) {
	params := db.FindByProviderParams{
		Provider:   provider,
		ProviderID: providerID,
	}

	row, err := O.queries.FindByProvider(ctx, params)
	if err != nil {
		return domain.OAuthCredential{}, err
	}

	return domain.OAuthCredential{
		ID:         row.ID,
		UserID:     row.UserID,
		Provider:   row.Provider,
		ProviderID: row.ProviderID,
	}, nil
}
