-- name: CreateToken :one
INSERT INTO refresh_tokens(user_id, token_hash, expires_at)
VALUES ($1, $2, $3)
RETURNING id;

-- name: FindByUserID :one
SELECT * FROM refresh_tokens
WHERE user_id = $1 and is_revoked = FALSE;

-- name: FindByTokenHash :one
SELECT * FROM refresh_tokens
WHERE token_hash = $1 and is_revoked = FALSE;

-- name: DeleteByUserID :exec
DELETE FROM refresh_tokens
WHERE user_id = $1;