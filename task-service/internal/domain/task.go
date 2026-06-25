package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Title       string
	Status      TaskStatus
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskStatus string

const (
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusDone       TaskStatus = "done"
	TaskStatusCancelled  TaskStatus = "cancelled"
)

type ListTasksParams struct {
	UserID   uuid.UUID
	Status   *TaskStatus
	Cursor   *time.Time
	PageSize int32
}
