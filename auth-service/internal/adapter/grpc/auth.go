package grpc

import (
	"context"
	"task-tracker/auth-service/internal/domain"
	"task-tracker/auth-service/internal/port/in"
	"task-tracker/gen/proto/auth"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthServer struct {
	service in.AuthService
	auth.UnimplementedAuthServiceServer
}

func NewAuthServer(service in.AuthService) *AuthServer {
	return &AuthServer{
		service: service,
	}
}

func (a AuthServer) LoginByEmail(ctx context.Context, request *auth.LoginByEmailRequest) (*auth.LoginByEmailResponse, error) {
	email := request.Email
	password := request.Password

	response, err := a.service.LoginByEmail(ctx, email, password)
	if err != nil {
		return nil, err
	}

	return &auth.LoginByEmailResponse{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	}, nil
}

func (a AuthServer) RefreshToken(ctx context.Context, request *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthServer) ValidateToken(ctx context.Context, request *auth.ValidateTokenRequest) (*auth.ValidateTokenResponse, error) {
	accessToken := request.AccessToken

	userID, err := a.service.ValidateToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	return &auth.ValidateTokenResponse{UserId: userID.String()}, nil
}

func (a AuthServer) Logout(ctx context.Context, request *auth.LogoutRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
