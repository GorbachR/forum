-- name: CreateRole :one
INSERT INTO roles (name, description) VALUES ($1, $2) RETURNING *;

-- name: ListRoles :many
SELECT * FROM roles;

-- name: GetRoles :one
SELECT * FROM roles WHERE role_id = $1;

-- name: UpdateRole :exec
UPDATE roles SET name = $2, description = $3 WHERE role_id = $1;

-- name: DeleteRole :exec
DELETE FROM roles WHERE role_id = $1;
