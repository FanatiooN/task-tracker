-- name: SaveContact :one
INSERT INTO user_contacts(user_id, provider, contact)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetContact :one
SELECT contact from user_contacts
WHERE user_id = $1 and provider = $2;