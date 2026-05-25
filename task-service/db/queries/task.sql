-- name: CreateTask :one
INSERT INTO tasks (user_id, title, description, status)
VALUES ($1, $2, $3, $4)
RETURNING id, title, description, status;

-- name: GetTask :one
SELECT * FROM tasks
WHERE deleted_at IS NULL AND id = $1;

-- name: GetTasksByUserID :many
SELECT * FROM tasks
WHERE deleted_at IS NULL AND user_id = $1;

-- name: GetTasksInProgressByUserID :many
SELECT * FROM tasks
WHERE deleted_at IS NULL AND user_id = $1 AND status = 'in_progress';

-- name: GetCompletedTasksTodayByUserID :many
SELECT * FROM tasks
WHERE deleted_at IS NULL AND user_id = $1 AND status = 'done'
AND updated_at >= date_trunc('day', now() AT TIME ZONE 'UTC');

-- name: UpdateTask :one
UPDATE tasks
SET title = $2, description = $3, status = $4, updated_at = now()
WHERE deleted_at IS NULL AND id = $1
RETURNING id, title, description, status;

-- name: DeleteTasks :many
UPDATE tasks
SET deleted_at = now(), updated_at = now()
WHERE deleted_at IS NULL and id = ANY($1::uuid[])
RETURNING id, title, description, status;