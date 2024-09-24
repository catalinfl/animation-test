package routes

import (
	"github.com/catalinfl/blogapi/handlers"
	"github.com/go-chi/chi/v5"
)

func AuthorsRoutes(r chi.Router, handler *handlers.Handler) {

	r.Get("/author/{id}", handler.GetAuthor)
	r.Post("/author", handler.CreateAuthor)
	r.Get("/authors", handler.GetAuthors)

}
