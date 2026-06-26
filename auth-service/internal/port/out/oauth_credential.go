package out

import (
	"context"
	"task-tracker/auth-service/internal/domain"
)

type OAuthCredentialRepository interface {
	Save(ctx context.Context, credential domain.OAuthCredential) error
	FindByProvider(ctx context.Context, provider, providerID string) (domain.OAuthCredential, error)
}
