// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: posts.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (content, user_id, topic_id) VALUES ($1, $2, $3) RETURNING post_id, content, user_id, topic_id, created_at
`

type CreatePostParams struct {
	Content string
	UserID  sql.NullInt32
	TopicID sql.NullInt32
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost, arg.Content, arg.UserID, arg.TopicID)
	var i Post
	err := row.Scan(
		&i.PostID,
		&i.Content,
		&i.UserID,
		&i.TopicID,
		&i.CreatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts WHERE post_id = $1
`

func (q *Queries) DeletePost(ctx context.Context, postID int32) error {
	_, err := q.db.ExecContext(ctx, deletePost, postID)
	return err
}

const getPosts = `-- name: GetPosts :one
SELECT post_id, content, user_id, topic_id, created_at FROM posts WHERE post_id = $1
`

func (q *Queries) GetPosts(ctx context.Context, postID int32) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPosts, postID)
	var i Post
	err := row.Scan(
		&i.PostID,
		&i.Content,
		&i.UserID,
		&i.TopicID,
		&i.CreatedAt,
	)
	return i, err
}

const listPosts = `-- name: ListPosts :many
SELECT post_id, content, user_id, topic_id, created_at FROM posts
`

func (q *Queries) ListPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.PostID,
			&i.Content,
			&i.UserID,
			&i.TopicID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePost = `-- name: UpdatePost :exec
UPDATE posts SET content = $2 WHERE post_id = $1
`

type UpdatePostParams struct {
	PostID  int32
	Content string
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) error {
	_, err := q.db.ExecContext(ctx, updatePost, arg.PostID, arg.Content)
	return err
}
