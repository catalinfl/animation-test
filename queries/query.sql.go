// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package queries

import (
	"context"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (name, email) VALUES ($1, $2)
RETURNING id, name, email, created_at
`

type CreateAuthorParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRow(ctx, createAuthor, arg.Name, arg.Email)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}

const createComment = `-- name: CreateComment :one
INSERT INTO comments (content, author_id, post_id) VALUES
($1, $2, $3) RETURNING id, content, author_id, post_id
`

type CreateCommentParams struct {
	Content  string `json:"content"`
	AuthorID int64  `json:"author_id"`
	PostID   int64  `json:"post_id"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRow(ctx, createComment, arg.Content, arg.AuthorID, arg.PostID)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.AuthorID,
		&i.PostID,
	)
	return i, err
}

const createPost = `-- name: CreatePost :one
INSERT INTO posts (title, content, author_id)
VALUES ($1, $2, $3) 
RETURNING id, title, content, created_at, updated_at, author_id
`

type CreatePostParams struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int64  `json:"author_id"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, createPost, arg.Title, arg.Content, arg.AuthorID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AuthorID,
	)
	return i, err
}

const getAllAuthors = `-- name: GetAllAuthors :many
SELECT id, name, email, created_at FROM authors
`

func (q *Queries) GetAllAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.Query(ctx, getAllAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, email, created_at FROM authors WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRow(ctx, getAuthor, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}

const getCommentsForAuthor = `-- name: GetCommentsForAuthor :many
SELECT id, content, author_id, post_id FROM comments WHERE
author_id = $1
`

func (q *Queries) GetCommentsForAuthor(ctx context.Context, authorID int64) ([]Comment, error) {
	rows, err := q.db.Query(ctx, getCommentsForAuthor, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.AuthorID,
			&i.PostID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCommentsForPost = `-- name: GetCommentsForPost :many
SELECT id, content, author_id, post_id FROM comments WHERE
post_id = $1
`

func (q *Queries) GetCommentsForPost(ctx context.Context, postID int64) ([]Comment, error) {
	rows, err := q.db.Query(ctx, getCommentsForPost, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.AuthorID,
			&i.PostID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPost = `-- name: GetPost :one
SELECT id, title, content, created_at, updated_at, author_id FROM posts WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, id int64) (Post, error) {
	row := q.db.QueryRow(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AuthorID,
	)
	return i, err
}

const getPosts = `-- name: GetPosts :many
SELECT id, title, content, created_at, updated_at, author_id FROM posts
`

func (q *Queries) GetPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.Query(ctx, getPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AuthorID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
