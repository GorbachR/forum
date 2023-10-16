-- name: CreateSession :one
INSERT INTO sessions (user_id, session_key, expires_at) VALUES ($1, $2, $3) RETURNING *;

-- name: ListSessions :many
SELECT * FROM sessions;

-- name: GetSesssion :one
SELECT * FROM sessions WHERE user_id = $1;

-- name: UpdateSession :exec
UPDATE sessions SET session_key = $2 WHERE user_id = $1;

-- name: DeleteSession :exec
DELETE FROM sessions WHERE user_id = $1;
