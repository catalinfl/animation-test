package routes

import (
	"net/http"

	"github.com/catalinfl/blogapi/queries"
	"github.com/go-chi/chi/v5"
)

func CommentsRoutes(r chi.Router, q *queries.Repo) {
	r.Get("/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("comments"))
	})
}
