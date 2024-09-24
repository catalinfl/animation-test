package main

import (
	"context"
	"log"
	"net/http"

	"github.com/catalinfl/blogapi/database"
	"github.com/catalinfl/blogapi/handlers"
	"github.com/catalinfl/blogapi/queries"
	"github.com/catalinfl/blogapi/routes"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	conn := database.ConnectDatabase()
	defer conn.Close()

	err := database.CreateTables(conn, context.Background())

	if err != nil {
		panic(err)
	}

	queries := queries.NewRepo(conn)
	handler := handlers.NewHandler(queries)

	routes.AuthorsRoutes(r, handler)
	routes.CommentsRoutes(r, handler)
	routes.PostsRoutes(r, handler)

	log.Printf("Starting server on port 8080")

	http.ListenAndServe(":8080", r)
}
