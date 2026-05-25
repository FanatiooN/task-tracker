-- name: CreateUser :one
INSERT INTO users (name, email)
VALUES ($1, $2)
RETURNING name, email;

-- name: GetUser :one
SELECT * FROM users
WHERE deleted_at IS NULL and id = $1;

-- name: UpdateUser :one
UPDATE users
SET name = $2, email = $3, updated_at = now()
WHERE deleted_at IS NULL and id = $1
RETURNING name, email;

-- name: DeleteUser :one
UPDATE users
SET deleted_at = now(), updated_at = now()
WHERE deleted_at IS NULL and id = $1
RETURNING name, email;