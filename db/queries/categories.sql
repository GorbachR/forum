-- name: CreateCategorie :one
INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING *;

-- name: ListCategories :many
SELECT * FROM categories;

-- name: GetCategorie :one
SELECT * FROM categories WHERE category_id = $1;

-- name: UpdateCategorie :exec
UPDATE categories SET name = $2, description = $3 WHERE category_id = $1;

-- name: DeleteCategorie :exec
DELETE FROM categories WHERE category_id = $1;

