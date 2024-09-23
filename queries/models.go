// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package queries

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Author struct {
	ID        int64            `json:"id"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type Comment struct {
	ID       int64  `json:"id"`
	Content  string `json:"content"`
	AuthorID int64  `json:"author_id"`
}

type Post struct {
	ID        int64            `json:"id"`
	Title     string           `json:"title"`
	Content   string           `json:"content"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	AuthorID  int64            `json:"author_id"`
}
