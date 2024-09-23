package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/catalinfl/blogapi/queries"
	"github.com/go-chi/chi/v5"
)

func AuthorsRoutes(r chi.Router, q *queries.Repo) {
	r.Get("/author/{id}", func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		id := chi.URLParam(r, "id")
		idNumber, err := strconv.Atoi(id)

		if err != nil {
			http.Error(w, "Invalid author id", http.StatusBadRequest)
			return
		}

		author, err := q.GetAuthor(ctx, int64(idNumber))

		if err != nil {
			http.Error(w, "Author not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(author)

	})

	r.Get("/authors", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		authors, err := q.GetAllAuthors(ctx)

		if err != nil {
			http.Error(w, "Error fetching authors", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(authors)
	})

	r.Post("/author", func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		var author queries.CreateAuthorParams

		err := json.NewDecoder(r.Body).Decode(&author)

		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		newAuthor, err := q.CreateAuthor(ctx, author)

		if err != nil {
			log.Fatalf("Error creating author: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newAuthor)
	})
}
