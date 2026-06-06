package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
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

func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {

	ownerId := strings.TrimSpace(r.URL.Query().Get("ownerId"))
	taskStatus := strings.TrimSpace(r.URL.Query().Get("taskStatus"))
	pageSize := strings.TrimSpace(r.URL.Query().Get("pageSize"))
	pageToken := strings.TrimSpace(r.URL.Query().Get("pageToken"))

	if ownerId == "" {
		writeError(w, http.StatusBadRequest, "missing required parameters")
		return
	}
	var size int32
	if pageSize == "" {
		size = 20
	} else {
		parsedSize, err := strconv.ParseInt(pageSize, 10, 32)
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid page size")
			return
		}
		size = int32(parsedSize)
	}

	var status *taskpb.TaskStatus
	if taskStatus != "" {
		value := taskpb.TaskStatus_value[taskStatus]
		if value == 0 {
			writeError(w, http.StatusBadRequest, "invalid task status")
			return
		}
		v := taskpb.TaskStatus(value)
		status = &v
	}

	response, err := h.client.ListTasks(r.Context(), &taskpb.ListTasksRequest{
		UserId:    ownerId,
		Status:    status,
		PageSize:  size,
		PageToken: pageToken,
	})
	if err != nil {
		//writeGRPCError(w, err) // TODO
		return
	}

	writeJSON(w, http.StatusOK, response)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	var req struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		TaskStatus  *string `json:"taskStatus"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	var status *taskpb.TaskStatus
	if req.TaskStatus != nil {
		value := taskpb.TaskStatus_value[*req.TaskStatus]
		if value == 0 {
			writeError(w, http.StatusBadRequest, "invalid task status")
			return
		}
		v := taskpb.TaskStatus(value)
		status = &v
	}

	response, err := h.client.UpdateTask(r.Context(), &taskpb.UpdateTaskRequest{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
		Status:      status,
	})
	if err != nil {
		//writeGRPCError(w, err) // TODO
		return
	}

	writeJSON(w, http.StatusOK, response.Task)
}

func (h *TaskHandler) DeleteTasks(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Ids []string `json:"ids"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	_, err := h.client.DeleteTasks(r.Context(), &taskpb.DeleteTasksRequest{
		Ids: req.Ids,
	})
	if err != nil {
		//writeGRPCError(w, err) // TODO
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
