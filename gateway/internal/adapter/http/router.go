package http

import "net/http"

type Handlers struct {
	User *UserHandler
	Task *TaskHandler
	Auth *AuthHandler
}

func NewRouter(h Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", h.User.CreateUser)
	mux.HandleFunc("GET /users/{id}", h.User.GetUser)
	mux.HandleFunc("PUT /users/{id}", h.User.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", h.User.DeleteUser)

	mux.HandleFunc("POST /login", h.Auth.Login)
	mux.HandleFunc("POST /refresh", h.Auth.Refresh)
	mux.HandleFunc("POST /logout", h.Auth.Logout)

	mux.HandleFunc("POST /tasks", h.Task.CreateTask)
	mux.HandleFunc("GET /tasks/{id}", h.Task.GetTask)
	mux.HandleFunc("GET /tasks", h.Task.ListTasks)
	mux.HandleFunc("PUT /tasks/{id}", h.Task.UpdateTask)
	mux.HandleFunc("DELETE /tasks/{id}", h.Task.DeleteTasks)

	return mux
}
