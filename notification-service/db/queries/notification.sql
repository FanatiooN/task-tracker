-- name: SaveNotification :one
INSERT INTO notifications(user_id, type, provider)
VALUES ($1, $2, $3)
RETURNING *;
