-- name: CreateCredentials :one
INSERT INTO credentials (user_id, email, password_hash)
VALUES ($1, $2, $3)
RETURNING user_id;

-- name: FindByEmail :one
SELECT * FROM credentials
WHERE email = $1;