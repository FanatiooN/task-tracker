package grpc

import (
	"context"
	userpb "task-tracker/gen/proto/user"
	"task-tracker/user-service/internal/domain"
	"task-tracker/user-service/internal/port/in"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserServer struct {
	service in.UserService
}

func (u UserServer) CreateUser(ctx context.Context, request *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user := domain.User{
		Name:  request.Name,
		Email: request.Email,
	}

	response, err := u.service.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	createResponse := userpb.User{
		Id:        response.ID.String(),
		Name:      response.Name,
		Email:     response.Email,
		CreatedAt: timestamppb.New(response.CreatedAt),
		UpdatedAt: timestamppb.New(response.UpdatedAt),
	}
	return &userpb.CreateUserResponse{User: &createResponse}, nil
}

func (u UserServer) GetUser(ctx context.Context, request *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServer) UpdateUser(ctx context.Context, request *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServer) DeleteUser(ctx context.Context, request *userpb.DeleteUserRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServer) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewUserServer(service in.UserService) *UserServer {
	return &UserServer{service: service}
}
