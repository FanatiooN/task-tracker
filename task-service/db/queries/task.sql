-- name: CreateTask :one
INSERT INTO tasks (user_id, title, description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks
WHERE deleted_at IS NULL AND id = $1;

-- name: ListTasks :many
SELECT * FROM tasks
WHERE deleted_at IS NULL AND user_id = $1
AND ($2::task_status IS NULL OR status = $2)
AND ($3::timestamptz IS NULL OR created_at < $3)
ORDER BY created_at DESC
LIMIT $4;

-- name: UpdateTask :one
UPDATE tasks
SET title = $2, description = $3, status = $4, updated_at = now()
WHERE deleted_at IS NULL AND id = $1
RETURNING *;

-- name: DeleteTasks :many
UPDATE tasks
SET deleted_at = now(), updated_at = now()
WHERE deleted_at IS NULL and id = ANY($1::uuid[])
RETURNING *;