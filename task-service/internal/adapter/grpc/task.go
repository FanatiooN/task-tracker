package grpc

import (
	"context"
	taskpb "task-tracker/gen/proto/task"
	"task-tracker/task-service/internal/domain"
	"task-tracker/task-service/internal/port/in"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskServer struct {
	taskpb.UnimplementedTaskServiceServer

	service in.TaskService
}

func NewTaskServer(service in.TaskService) *TaskServer {
	return &TaskServer{service: service}
}

func (t TaskServer) CreateTask(ctx context.Context, request *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}

	var description string
	if request.Description != nil {
		description = *request.Description
	}

	task := domain.Task{
		UserID:      userID,
		Title:       request.Title,
		Description: description,
	}

	response, err := t.service.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}

	return &taskpb.CreateTaskResponse{
		Task: &taskpb.Task{
			Id:          response.ID.String(),
			Title:       response.Title,
			Description: &response.Description,
			Status:      convertDomainToProtoStatus(response.Status),
			UserId:      response.UserID.String(),
			CreatedAt:   timestamppb.New(response.CreatedAt),
			UpdatedAt:   timestamppb.New(response.UpdatedAt),
		},
	}, nil
}

func (t TaskServer) GetTask(ctx context.Context, request *taskpb.GetTaskRequest) (*taskpb.GetTaskResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}

	response, err := t.service.GetTask(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	return &taskpb.GetTaskResponse{
		Task: &taskpb.Task{
			Id:          response.ID.String(),
			Title:       response.Title,
			Description: &response.Description,
			Status:      convertDomainToProtoStatus(response.Status),
			UserId:      response.UserID.String(),
			CreatedAt:   timestamppb.New(response.CreatedAt),
			UpdatedAt:   timestamppb.New(response.UpdatedAt),
		},
	}, nil
}

func (t TaskServer) ListTasks(ctx context.Context, request *taskpb.ListTasksRequest) (*taskpb.ListTasksResponse, error) {
	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}

	params := domain.ListTasksParams{
		UserID:   userID,
		Status:   convertProtoToDomainStatus(request.Status),
		PageSize: request.PageSize,
	}

	response, cursor, err := t.service.ListTasks(ctx, request.PageToken, params)
	if err != nil {
		return nil, err
	}

	var tasks []*taskpb.Task

	for _, r := range response {
		tasks = append(tasks, &taskpb.Task{
			Id:          r.ID.String(),
			Title:       r.Title,
			Description: &r.Description,
			Status:      convertDomainToProtoStatus(r.Status),
			UserId:      r.UserID.String(),
			CreatedAt:   timestamppb.New(r.CreatedAt),
			UpdatedAt:   timestamppb.New(r.UpdatedAt),
		})
	}

	return &taskpb.ListTasksResponse{
		Tasks:         tasks,
		NextPageToken: cursor,
	}, nil
}

func (t TaskServer) UpdateTask(ctx context.Context, request *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	taskID, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}

	var description string
	if request.Description != nil {
		description = *request.Description
	}

	var title string
	if request.Title != nil {
		title = *request.Title
	}

	var status domain.TaskStatus
	if request.Status != nil {
		status = *convertProtoToDomainStatus(request.Status)
	}

	response, err := t.service.UpdateTask(ctx, domain.Task{
		ID:          taskID,
		Title:       title,
		Status:      status,
		Description: description,
	}, userID)
	if err != nil {
		return nil, err
	}

	return &taskpb.UpdateTaskResponse{
		Task: &taskpb.Task{
			Id:          response.ID.String(),
			Title:       response.Title,
			Description: &response.Description,
			Status:      convertDomainToProtoStatus(response.Status),
			UserId:      response.UserID.String(),
			CreatedAt:   timestamppb.New(response.CreatedAt),
			UpdatedAt:   timestamppb.New(response.UpdatedAt),
		}}, nil
}

func (t TaskServer) DeleteTasks(ctx context.Context, request *taskpb.DeleteTasksRequest) (*emptypb.Empty, error) {
	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}

	ids := make([]uuid.UUID, 0, len(request.Ids))
	for _, id := range request.Ids {
		parsedId, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}

		ids = append(ids, parsedId)
	}

	err = t.service.DeleteTasks(ctx, ids, userID)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (t TaskServer) SendDailyReport(ctx context.Context, request *emptypb.Empty) (*emptypb.Empty, error) {
	err := t.service.SendDailyReport(ctx)

	return &emptypb.Empty{}, err
}

func convertDomainToProtoStatus(status domain.TaskStatus) taskpb.TaskStatus {
	switch status {
	case domain.TaskStatusInProgress:
		return taskpb.TaskStatus_TASK_STATUS_IN_PROGRESS
	case domain.TaskStatusDone:
		return taskpb.TaskStatus_TASK_STATUS_DONE
	case domain.TaskStatusCancelled:
		return taskpb.TaskStatus_TASK_STATUS_CANCELLED
	}

	return taskpb.TaskStatus_TASK_STATUS_UNSPECIFIED
}

func convertProtoToDomainStatus(status *taskpb.TaskStatus) *domain.TaskStatus {
	if status == nil {
		return nil
	}

	var s domain.TaskStatus

	switch *status {
	case taskpb.TaskStatus_TASK_STATUS_IN_PROGRESS:
		s = domain.TaskStatusInProgress
	case taskpb.TaskStatus_TASK_STATUS_DONE:
		s = domain.TaskStatusDone
	case taskpb.TaskStatus_TASK_STATUS_CANCELLED:
		s = domain.TaskStatusCancelled
	case taskpb.TaskStatus_TASK_STATUS_UNSPECIFIED:
		return nil
	}

	return &s
}
