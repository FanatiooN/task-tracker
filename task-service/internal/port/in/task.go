package in

import (
	"context"
	"task-tracker/task-service/internal/domain"

	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTask(ctx context.Context, id uuid.UUID) (domain.Task, error)
	ListTasks(ctx context.Context, pageToken string, params domain.ListTasksParams) ([]domain.Task, string, error)
	UpdateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	DeleteTasks(ctx context.Context, id []uuid.UUID) error
}
