package out

import (
	"context"
	"task-tracker/auth-service/internal/domain"
)

type OAuthProvider interface {
	GetUserInfo(ctx context.Context, provider, code, redirectURI string) (domain.OAuthUserInfo, error)
}
