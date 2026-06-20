package service

import (
	"context"
	"task-tracker/auth-service/internal/domain"
	"task-tracker/auth-service/internal/port/out"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	credentials out.CredentialRepository
	tokens      out.TokenRepository
	jwtSecret   string
	accessTTL   time.Duration
	refreshTTL  time.Duration
}

func NewAuthService(credentials out.CredentialRepository, tokens out.TokenRepository, jwtSecret string, accessTTL time.Duration, refreshTTL time.Duration) *AuthService {
	return &AuthService{
		credentials: credentials,
		tokens:      tokens,
		jwtSecret:   jwtSecret,
		accessTTL:   accessTTL,
		refreshTTL:  refreshTTL,
	}
}

func (a AuthService) LoginByEmail(ctx context.Context, email, password string) (domain.Tokens, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthService) RefreshToken(ctx context.Context, refreshToken string) (domain.Tokens, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthService) Logout(ctx context.Context, refreshToken string) error {
	//TODO implement me
	panic("implement me")
}

func (a AuthService) ValidateToken(ctx context.Context, accessToken string) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}
