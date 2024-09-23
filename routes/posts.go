package routes

import (
	"net/http"

	"github.com/catalinfl/blogapi/queries"
	"github.com/go-chi/chi/v5"
)

func PostsRoutes(r chi.Router, q *queries.Repo) {
	r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("posts"))
	})
}
