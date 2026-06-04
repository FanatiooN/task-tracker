package http

import (
	"encoding/json"
	"net/http"
	taskpb "task-tracker/gen/proto/task"
)

type TaskHandler struct {
	client taskpb.TaskServiceClient
}

func NewTaskHandler(client taskpb.TaskServiceClient) *TaskHandler {
	return &TaskHandler{client: client}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserId      string  `json:"userId"`
		Title       string  `json:"title"`
		Description *string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	response, err := h.client.CreateTask(r.Context(), &taskpb.CreateTaskRequest{
		UserId:      req.UserId,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		//writeGRPCError(w, err) // TODO
		return
	}

	writeJSON(w, http.StatusCreated, response.Task)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	response, err := h.client.GetTask(r.Context(), &taskpb.GetTaskRequest{
		Id: id,
	})
	if err != nil {
		//writeGRPCError(w, err) // TODO
		return
	}

	writeJSON(w, http.StatusOK, response.Task)
}
