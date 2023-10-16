-- name: CreateTopic :one
INSERT INTO topics (title, content, user_id, category_id) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: ListTopics :many
SELECT * FROM topics;

-- name: GetTopics :one
SELECT * FROM topics WHERE topic_id = $1;

-- name: UpdateTopic :exec
UPDATE topics SET title = $2, content = $3 WHERE topic_id = $1;

-- name: DeleteTopic :exec
DELETE FROM topics WHERE topic_id = $1;

