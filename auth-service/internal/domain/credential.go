package domain

import (
	"github.com/google/uuid"
)

type Credential struct {
	UserID       uuid.UUID
	email        string
	PasswordHash string
}
