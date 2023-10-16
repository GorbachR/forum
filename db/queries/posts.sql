-- name: CreatePost :one
INSERT INTO posts (content, user_id, topic_id) VALUES ($1, $2, $3) RETURNING *;

-- name: ListPosts :many
SELECT * FROM posts;

-- name: GetPosts :one
SELECT * FROM posts WHERE post_id = $1;

-- name: UpdatePost :exec
UPDATE posts SET content = $2 WHERE post_id = $1;

-- name: DeletePost :exec
DELETE FROM posts WHERE post_id = $1;
