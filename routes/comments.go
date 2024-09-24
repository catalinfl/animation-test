package routes

import (
	"github.com/catalinfl/blogapi/handlers"
	"github.com/go-chi/chi/v5"
)

func CommentsRoutes(r chi.Router, h *handlers.Handler) {
	r.Get("/comments-post/{postId}", h.GetCommentsForPost)
	r.Get("/comments-author/{authorId}", h.GetCommentsForAuthor)
	r.Post("/comments", h.CreateComment)
}
