package http

import (
	"encoding/json"
	"net/http"
	authpb "task-tracker/gen/proto/auth"
	userpb "task-tracker/gen/proto/user"
)

type AuthHandler struct {
	client     authpb.AuthServiceClient
	userClient userpb.UserServiceClient
}

func NewAuthHandler(client authpb.AuthServiceClient, userClient userpb.UserServiceClient) *AuthHandler {
	return &AuthHandler{client: client, userClient: userClient}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	response, err := h.client.LoginByEmail(r.Context(), &authpb.LoginByEmailRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"access_token":  response.AccessToken,
		"refresh_token": response.RefreshToken,
	})
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	response, err := h.client.RefreshToken(r.Context(), &authpb.RefreshTokenRequest{
		RefreshToken: req.Token,
	})
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"access_token":  response.AccessToken,
		"refresh_token": response.RefreshToken,
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	_, err := h.client.Logout(r.Context(), &authpb.LogoutRequest{
		RefreshToken: req.Token,
	})
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	userResponse, err := h.userClient.CreateUser(r.Context(), &userpb.CreateUserRequest{
		Name:  req.Name,
		Email: &req.Email,
	})
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	response, err := h.client.RegisterByEmail(r.Context(), &authpb.RegisterByEmailRequest{
		Email:    req.Email,
		Password: req.Password,
		UserId:   userResponse.User.Id,
	})
	if err != nil {
		writeGRPCError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"access_token":  response.AccessToken,
		"refresh_token": response.RefreshToken,
	})
}
