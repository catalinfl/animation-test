package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTables(p *pgxpool.Pool, ctx context.Context) error {

	authorTable := `
	CREATE TABLE IF NOT EXISTS authors (
		id BIGSERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW()
	);

	`

	postTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id BIGSERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		author_id BIGINT NOT NULL,
		FOREIGN KEY (author_id) REFERENCES authors(id)
	);
	`

	comments := `
		CREATE TABLE IF NOT EXISTS comments (
		id BIGSERIAL PRIMARY KEY,
		content TEXT NOT NULL,
		author_id BIGINT NOT NULL,
		post_id BIGINT NOT NULL,
		FOREIGN KEY (author_id) REFERENCES authors(id),
		FOREIGN KEY (post_id) REFERENCES posts(id)
	)
	`

	_, err := p.Exec(ctx, authorTable)

	if err != nil {
		return err
	}

	_, err = p.Exec(ctx, postTable)

	if err != nil {
		return err
	}

	_, err = p.Exec(ctx, comments)

	if err != nil {
		return err
	}

	return nil
}
