-- name: CreateUser :one
INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING *;

-- name: ListUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE user_id = $1;

-- name: UpdateUser :exec
UPDATE users SET username = $2, email = $3, password_hash = $4, dark_theme = $5 WHERE user_id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = $1;
