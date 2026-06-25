package service

import (
	"context"
	"errors"
	"strings"
	"task-tracker/task-service/internal/domain"
	"task-tracker/task-service/internal/port/out"

	"github.com/google/uuid"
)

type TaskService struct {
	repository out.TaskRepository
}

func NewTaskService(repository out.TaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

func (t TaskService) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	task.Title = strings.TrimSpace(task.Title)
	if task.Title == "" || len(task.Title) > domain.MaxTitleLength {
		return domain.Task{}, errors.New("wrong title length")
	}

	task.Description = strings.TrimSpace(task.Description)
	if len(task.Description) > domain.MaxDescriptionLength {
		return domain.Task{}, errors.New("wrong description length")
	}

	createdTask, err := t.repository.Save(ctx, task)
	if err != nil {
		return domain.Task{}, err
	}

	return createdTask, nil
}

func (t TaskService) GetTask(ctx context.Context, id uuid.UUID) (domain.Task, error) {
	task, err := t.repository.FindByID(ctx, id)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (t TaskService) ListTasks(ctx context.Context, pageToken string, params domain.ListTasksParams) ([]domain.Task, string, error) {
	if params.PageSize > domain.MaxPageSize {
		return nil, "", errors.New("page size too large")
	}

	if pageToken != "" {
		cursor, err := decodeToken(pageToken)
		if err != nil {
			return nil, "", err
		}
		params.Cursor = cursor
	}

	tasks, err := t.repository.List(ctx, params)
	if err != nil {
		return nil, "", err
	}

	var cursor string

	if int32(len(tasks)) == params.PageSize {
		cursor = encodeToken(tasks[len(tasks)-1].CreatedAt)
	}

	return tasks, cursor, nil

}

func (t TaskService) UpdateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	existedTask, err := t.repository.FindByID(ctx, task.ID)
	if err != nil {
		return domain.Task{}, err
	}

	task.Title = strings.TrimSpace(task.Title)
	if task.Title != "" {
		if len(task.Title) <= domain.MaxTitleLength {
			existedTask.Title = task.Title
		} else {
			return domain.Task{}, errors.New("wrong title length")
		}
	}

	task.Description = strings.TrimSpace(task.Description)
	if task.Description != "" {
		if len(task.Description) <= domain.MaxDescriptionLength {
			existedTask.Description = task.Description
		} else {
			return domain.Task{}, errors.New("wrong description length")
		}
	}

	if task.Status != "" {
		existedTask.Status = task.Status
	}

	updatedTask, err := t.repository.Update(ctx, existedTask)
	if err != nil {
		return domain.Task{}, err
	}

	return updatedTask, nil
}

func (t TaskService) DeleteTasks(ctx context.Context, id []uuid.UUID) error {
	if len(id) == 0 {
		return errors.New("empty task id list")
	}

	err := t.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
