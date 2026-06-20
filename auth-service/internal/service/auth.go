package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"task-tracker/auth-service/internal/domain"
	"task-tracker/auth-service/internal/port/out"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	credentials out.CredentialRepository
	tokens      out.TokenRepository
	jwtSecret   string
	accessTTL   time.Duration
	refreshTTL  time.Duration
}

type AccessClaims struct {
	UserID uuid.UUID
	Email  string
	jwt.RegisteredClaims
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
	creds, err := a.credentials.FindByEmail(ctx, email)
	if err != nil {
		return domain.Tokens{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(creds.PasswordHash), []byte(password))
	if err != nil {
		return domain.Tokens{}, err
	}

	tokens, err := a.generateAccessToken(ctx, creds.UserID, creds.Email)
	if err != nil {
		return domain.Tokens{}, err
	}

	return tokens, nil
}

func (a AuthService) generateAccessToken(ctx context.Context, userID uuid.UUID, email string) (domain.Tokens, error) {
	now := time.Now()

	expiresAt := jwt.NewNumericDate(now.Add(a.accessTTL))
	claims := AccessClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(a.jwtSecret))
	if err != nil {
		return domain.Tokens{}, err
	}

	refreshToken := uuid.NewString()
	hash := sha256.Sum256([]byte(refreshToken))
	tokenHash := hex.EncodeToString(hash[:])

	err = a.tokens.Save(ctx, domain.RefreshToken{
		ID:        uuid.New(),
		UserID:    userID,
		TokenHash: tokenHash,
		ExpiresAt: expiresAt.Time,
		IsRevoked: false,
	})

	if err != nil {
		return domain.Tokens{}, err
	}

	return domain.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
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
	token, err := jwt.ParseWithClaims(accessToken, &AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.jwtSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	if !token.Valid {
		return uuid.Nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*AccessClaims)
	if !ok {
		return uuid.Nil, errors.New("invalid token")
	}

	return claims.UserID, nil
}
