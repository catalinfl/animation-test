-- name: GetAuthor :one
SELECT * FROM authors WHERE id = $1 LIMIT 1;

-- name: CreateAuthor :one
INSERT INTO authors (name, email) VALUES ($1, $2)
RETURNING *;

-- name: GetAllAuthors :many
SELECT * FROM authors;

-- name: CreatePost :one
INSERT INTO posts (title, content, author_id)
VALUES ($1, $2, $3) 
RETURNING *
;

-- name: GetPost :one
SELECT * FROM posts WHERE id = $1 LIMIT 1;

-- name: GetPosts :many
SELECT * FROM posts;