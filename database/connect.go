package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDatabase() *pgxpool.Pool {
	godotenv.Load()

	database_url := os.Getenv("DATABASE_URL")

	conn, err := pgxpool.New(context.Background(), database_url)

	if err != nil {
		fmt.Println("Error connecting to database: ", err)
	}

	return conn
}
