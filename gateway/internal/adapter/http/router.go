package http

import (
	"net/http"
	authpb "task-tracker/gen/proto/auth"
)

type Handlers struct {
	User  *UserHandler
	Task  *TaskHandler
	Auth  *AuthHandler
	OAuth *OAuthHandler
}

func NewRouter(h Handlers, authClient authpb.AuthServiceClient) http.Handler {
	mux := http.NewServeMux()

	auth := AuthMiddleware(authClient)

	mux.HandleFunc("POST /login", h.Auth.Login)
	mux.HandleFunc("POST /register", h.Auth.Register)
	mux.HandleFunc("POST /refresh", h.Auth.Refresh)

	mux.HandleFunc("GET /login/google", h.OAuth.LoginWithGoogle)
	mux.HandleFunc("GET /login/google/callback", h.OAuth.LoginCallbackWithGoogle)
	mux.HandleFunc("POST /login/telegram", h.OAuth.LoginWithTelegram)

	mux.Handle("GET /users/{id}", auth(http.HandlerFunc(h.User.GetUser)))
	mux.Handle("PUT /users/{id}", auth(http.HandlerFunc(h.User.UpdateUser)))
	mux.Handle("DELETE /users/{id}", auth(http.HandlerFunc(h.User.DeleteUser)))

	mux.Handle("POST /logout", auth(http.HandlerFunc(h.Auth.Logout)))

	mux.Handle("POST /tasks", auth(http.HandlerFunc(h.Task.CreateTask)))
	mux.Handle("GET /tasks/{id}", auth(http.HandlerFunc(h.Task.GetTask)))
	mux.Handle("GET /tasks", auth(http.HandlerFunc(h.Task.ListTasks)))
	mux.Handle("PUT /tasks/{id}", auth(http.HandlerFunc(h.Task.UpdateTask)))
	mux.Handle("DELETE /tasks", auth(http.HandlerFunc(h.Task.DeleteTasks)))

	return mux
}
