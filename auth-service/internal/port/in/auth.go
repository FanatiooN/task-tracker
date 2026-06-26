package in

import (
	"context"
	"task-tracker/auth-service/internal/domain"

	"github.com/google/uuid"
)

type AuthService interface {
	LoginByEmail(ctx context.Context, email, password string) (domain.Tokens, error)
	RegisterByEmail(ctx context.Context, name, email, password string) (domain.Tokens, error)
	RefreshToken(ctx context.Context, refreshToken string) (domain.Tokens, error)
	Logout(ctx context.Context, refreshToken string) error
	ValidateToken(ctx context.Context, accessToken string) (uuid.UUID, error)
	LoginByOAuth(ctx context.Context, provider, code, redirectURI string) (domain.Tokens, error)
}
