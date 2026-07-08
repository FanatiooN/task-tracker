package kafka

import "github.com/google/uuid"

type UserContactLinkedEvent struct {
	UserID   uuid.UUID `json:"user_id"`
	Provider string    `json:"provider"`
	Contact  string    `json:"contact"`
}
