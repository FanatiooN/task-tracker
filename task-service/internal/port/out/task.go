package out

import (
	"context"
	"task-tracker/task-service/internal/domain"

	"github.com/google/uuid"
)

type TaskRepository interface {
	Save(ctx context.Context, task domain.Task) (domain.Task, error)
	FindByID(ctx context.Context, id uuid.UUID) (domain.Task, error)
	List(ctx context.Context, params domain.ListTasksParams) ([]domain.Task, error)
	Update(ctx context.Context, task domain.Task) (domain.Task, error)
	Delete(ctx context.Context, id []uuid.UUID, userID uuid.UUID) error
}
