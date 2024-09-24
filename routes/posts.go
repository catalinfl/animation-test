package routes

import (
	"github.com/catalinfl/blogapi/handlers"
	"github.com/go-chi/chi/v5"
)

func PostsRoutes(r chi.Router, h *handlers.Handler) {
	r.Get("/posts", h.GetPosts)
	r.Get("/post/{id}", h.GetPost)
	r.Post("/post", h.CreatePost)
}
