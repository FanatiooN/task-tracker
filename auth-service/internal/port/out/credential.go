package out

import (
	"context"
	"task-tracker/auth-service/internal/domain"
)

type CredentialRepository interface {
	Save(ctx context.Context, credential domain.Credential) error
	FindByEmail(ctx context.Context, email string) (domain.Credential, error)
}
