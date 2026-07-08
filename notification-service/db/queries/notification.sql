-- name: SaveNotification :one
INSERT INTO notifications(user_id, type, provider)
VALUES ($1, $2, $3)
RETURNING *;


-- name: GetStatistics :many
SELECT type, provider, DATE(created_at) as date,
       COUNT(distinct user_id) AS unique_hits,
       COUNT(*) AS total_hits
FROM notifications
WHERE (sqlc.narg('provider')::text IS NULL OR provider = sqlc.narg('provider')::text)
AND (sqlc.narg('date')::timestamptz IS NULL OR DATE(created_at) = DATE(sqlc.narg('date')::timestamptz))
GROUP BY type, provider, DATE(created_at);