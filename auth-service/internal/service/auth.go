package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"task-tracker/auth-service/internal/domain"
	"task-tracker/auth-service/internal/port/out"
	userpb "task-tracker/gen/proto/user"
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
	userClient  userpb.UserServiceClient

	oauthProvider    out.OAuthProvider
	oauthCredentials out.OAuthCredentialRepository
}

type AccessClaims struct {
	jwt.RegisteredClaims

	UserID uuid.UUID
}

func NewAuthService(
	credentials out.CredentialRepository,
	tokens out.TokenRepository,
	jwtSecret string,
	accessTTL time.Duration,
	refreshTTL time.Duration,
	userClient userpb.UserServiceClient,
	oauthProvider out.OAuthProvider,
	oauthCredentials out.OAuthCredentialRepository,
) *AuthService {
	return &AuthService{
		credentials:      credentials,
		tokens:           tokens,
		jwtSecret:        jwtSecret,
		accessTTL:        accessTTL,
		refreshTTL:       refreshTTL,
		userClient:       userClient,
		oauthCredentials: oauthCredentials,
		oauthProvider:    oauthProvider,
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

	tokens, err := a.generateAccessToken(ctx, creds.UserID)
	if err != nil {
		return domain.Tokens{}, err
	}

	return tokens, nil
}

func (a AuthService) RegisterByEmail(ctx context.Context, name, email, password string) (domain.Tokens, error) {
	_, err := a.credentials.FindByEmail(ctx, email)
	if err == nil {
		return domain.Tokens{}, errors.New("email already registered")
	}

	createdUser, err := a.userClient.CreateUser(ctx, &userpb.CreateUserRequest{
		Name:  name,
		Email: &email,
	})
	if err != nil {
		return domain.Tokens{}, err
	}

	userID, err := uuid.Parse(createdUser.User.Id)
	if err != nil {
		return domain.Tokens{}, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.Tokens{}, err
	}

	err = a.credentials.Save(ctx, domain.Credential{
		UserID:       userID,
		Email:        email,
		PasswordHash: string(passwordHash),
	})
	if err != nil {
		return domain.Tokens{}, err
	}

	tokens, err := a.generateAccessToken(ctx, userID)
	if err != nil {
		return domain.Tokens{}, err
	}

	return tokens, nil
}

func (a AuthService) generateAccessToken(ctx context.Context, userID uuid.UUID) (domain.Tokens, error) {
	now := time.Now()

	expiresAt := jwt.NewNumericDate(now.Add(a.accessTTL))
	claims := AccessClaims{
		UserID: userID,
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

	_ = a.tokens.DeleteByUserID(ctx, userID)

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
	hash := sha256.Sum256([]byte(refreshToken))
	tokenHash := hex.EncodeToString(hash[:])

	token, err := a.tokens.FindByTokenHash(ctx, tokenHash)
	if err != nil {
		return domain.Tokens{}, err
	}

	if token.ExpiresAt.Before(time.Now()) {
		deleteErr := a.tokens.DeleteByUserID(ctx, token.UserID)
		if deleteErr != nil {
			return domain.Tokens{}, deleteErr
		}

		return domain.Tokens{}, errors.New("token expired")
	}

	err = a.tokens.DeleteByUserID(ctx, token.UserID)
	if err != nil {
		return domain.Tokens{}, err
	}

	tokens, err := a.generateAccessToken(ctx, token.UserID)
	if err != nil {
		return domain.Tokens{}, err
	}

	return tokens, nil
}

func (a AuthService) Logout(ctx context.Context, refreshToken string) error {
	hash := sha256.Sum256([]byte(refreshToken))
	tokenHash := hex.EncodeToString(hash[:])

	token, err := a.tokens.FindByTokenHash(ctx, tokenHash)
	if err != nil {
		return err
	}

	err = a.tokens.DeleteByUserID(ctx, token.UserID)
	if err != nil {
		return err
	}

	return nil
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

func (a AuthService) LoginByOAuth(ctx context.Context, provider, code, redirectURI string) (domain.Tokens, error) {
	userInfo, getErr := a.oauthProvider.GetUserInfo(ctx, provider, code, redirectURI)
	if getErr != nil {
		return domain.Tokens{}, getErr
	}

	creds, findErr := a.oauthCredentials.FindByProvider(ctx, provider, userInfo.ID)
	if findErr != nil {
		createdUser, err := a.userClient.CreateUser(ctx, &userpb.CreateUserRequest{
			Name:  userInfo.Name,
			Email: userInfo.Email,
		})
		if err != nil {
			return domain.Tokens{}, err
		}

		userID, err := uuid.Parse(createdUser.User.Id)
		if err != nil {
			return domain.Tokens{}, err
		}

		err = a.oauthCredentials.Save(ctx, domain.OAuthCredential{
			UserID:     userID,
			Provider:   provider,
			ProviderID: userInfo.ID,
		})
		if err != nil {
			return domain.Tokens{}, err
		}

		tokens, err := a.generateAccessToken(ctx, userID)
		if err != nil {
			return domain.Tokens{}, err
		}

		return tokens, nil
	}

	tokens, err := a.generateAccessToken(ctx, creds.UserID)
	if err != nil {
		return domain.Tokens{}, err
	}

	return tokens, nil
}
