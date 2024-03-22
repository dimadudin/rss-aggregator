-- name: CreateFollow :one
INSERT INTO follows (id, user_id, feed_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFollows :many
SELECT * FROM follows WHERE user_id = $1;

-- name: DeleteFollow :one
DELETE FROM follows WHERE id = $1 AND user_id = $2
RETURNING *;
