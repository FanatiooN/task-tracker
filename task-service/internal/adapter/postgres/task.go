package postgres

import (
	"context"
	"task-tracker/task-service/internal/db"
	"task-tracker/task-service/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskRepository struct {
	queries *db.Queries
}

func NewTaskRepository(queries *db.Queries) *TaskRepository {
	return &TaskRepository{queries: queries}
}

func (t TaskRepository) Save(ctx context.Context, task domain.Task) (domain.Task, error) {
	params := db.CreateTaskParams{
		UserID: task.UserID,
		Title:  task.Title,
	}

	if task.Description != "" {
		params.Description = pgtype.Text{
			String: task.Description,
			Valid:  true,
		}
	} else {
		params.Description = pgtype.Text{
			String: "",
			Valid:  false,
		}
	}

	row, err := t.queries.CreateTask(ctx, params)
	if err != nil {
		return domain.Task{}, err
	}

	return domain.Task{
		ID:          row.ID,
		UserID:      row.UserID,
		Title:       row.Title,
		Status:      domain.TaskStatus(row.Status),
		Description: row.Description.String,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}, nil
}

func (t TaskRepository) FindByID(ctx context.Context, id uuid.UUID) (domain.Task, error) {
	row, err := t.queries.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, err
	}

	return domain.Task{
		ID:          row.ID,
		UserID:      row.UserID,
		Title:       row.Title,
		Status:      domain.TaskStatus(row.Status),
		Description: row.Description.String,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}, nil
}

func (t TaskRepository) List(ctx context.Context, params domain.ListTasksParams) ([]domain.Task, error) {
	queryParams := db.ListTasksParams{
		UserID: params.UserID,
	}

	if params.Status != nil {
		status := *params.Status
		queryParams.Status = db.NullTaskStatus{TaskStatus: db.TaskStatus(status), Valid: true}
	} else {
		queryParams.Status = db.NullTaskStatus{Valid: false}
	}

	if params.Cursor != nil {
		queryParams.Cursor = pgtype.Timestamptz{Time: *params.Cursor, Valid: true}
	} else {
		queryParams.Cursor = pgtype.Timestamptz{Valid: false}
	}

	if params.PageSize > 0 {
		queryParams.Limit = params.PageSize
	} else {
		queryParams.Limit = 20
	}

	row, err := t.queries.ListTasks(ctx, queryParams)
	if err != nil {
		return nil, err
	}

	tasks := make([]domain.Task, 0, len(row))
	for idx := 0; idx < len(row); idx++ {
		tasks = append(tasks, domain.Task{
			ID:          row[idx].ID,
			UserID:      row[idx].UserID,
			Title:       row[idx].Title,
			Status:      domain.TaskStatus(row[idx].Status),
			Description: row[idx].Description.String,
			CreatedAt:   row[idx].CreatedAt,
			UpdatedAt:   row[idx].UpdatedAt,
		})
	}

	return tasks, nil
}

func (t TaskRepository) Update(ctx context.Context, task domain.Task) (domain.Task, error) {
	params := db.UpdateTaskParams{
		ID:     task.ID,
		Title:  task.Title,
		Status: db.TaskStatus(task.Status),
		Description: pgtype.Text{
			String: task.Description,
			Valid:  true,
		},
	}

	row, err := t.queries.UpdateTask(ctx, params)
	if err != nil {
		return domain.Task{}, err
	}

	return domain.Task{
		ID:          row.ID,
		UserID:      row.UserID,
		Title:       row.Title,
		Status:      domain.TaskStatus(row.Status),
		Description: row.Description.String,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}, nil
}

func (t TaskRepository) Delete(ctx context.Context, id []uuid.UUID) error {
	_, err := t.queries.DeleteTasks(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
