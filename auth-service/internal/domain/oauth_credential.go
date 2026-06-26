package domain

import "github.com/google/uuid"

type OAuthCredential struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	Provider   string
	ProviderID string
}

type OAuthUserInfo struct {
	ID         string
	Name       string
	Email      *string
	TGUsername *string
}
