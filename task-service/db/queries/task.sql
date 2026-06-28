-- name: CreateTask :one
INSERT INTO tasks (user_id, title, description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks
WHERE deleted_at IS NULL AND id = $1;

-- name: ListTasks :many
SELECT * FROM tasks
WHERE deleted_at IS NULL AND user_id = sqlc.arg('user_id')
AND (sqlc.narg('status')::task_status IS NULL OR status = sqlc.narg('status')::task_status)
AND (sqlc.narg('cursor')::timestamptz IS NULL OR created_at < sqlc.narg('cursor')::timestamptz)
ORDER BY created_at DESC
LIMIT sqlc.arg('limit');

-- name: UpdateTask :one
UPDATE tasks
SET title = $2, description = $3, status = $4, updated_at = now()
WHERE deleted_at IS NULL AND id = $1
RETURNING *;

-- name: DeleteTasks :many
UPDATE tasks
SET deleted_at = now(), updated_at = now()
WHERE deleted_at IS NULL AND id = ANY($1::uuid[]) AND user_id = $2
RETURNING *;