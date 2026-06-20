package out

import (
	"context"
	"task-tracker/auth-service/internal/domain"
)

type Credential interface {
	Save(ctx context.Context, credential domain.Credential)
	FindByEmail(ctx context.Context, email string)
}
