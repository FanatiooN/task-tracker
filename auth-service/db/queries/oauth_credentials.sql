-- name: CreateOAuthCredential :one
INSERT INTO oauth_credentials(user_id, provider, provider_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: FindByProvider :one
SELECT * FROM oauth_credentials
WHERE provider = $1 and provider_id = $2;