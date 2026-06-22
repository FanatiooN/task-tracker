package grpc

import (
	"context"
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

func (a AuthServer) RegisterByEmail(ctx context.Context, request *auth.RegisterByEmailRequest) (*auth.RegisterByEmailResponse, error) {
	email := request.Email
	password := request.Password

	userID, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}

	response, err := a.service.RegisterByEmail(ctx, userID, email, password)
	if err != nil {
		return nil, err
	}

	return &auth.RegisterByEmailResponse{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	}, nil
}

func (a AuthServer) RefreshToken(ctx context.Context, request *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	refreshToken := request.RefreshToken

	tokens, err := a.service.RefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	return &auth.RefreshTokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
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
	refreshToken := request.RefreshToken

	err := a.service.Logout(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
