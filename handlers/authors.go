package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/catalinfl/blogapi/queries"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Repo *queries.Repo
}

func NewHandler(repo *queries.Repo) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")

	idint, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	author, err := h.Repo.GetAuthor(ctx, int64(idint))
	if err != nil {
		http.Error(w, "Error getting author", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

func (h *Handler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authors, err := h.Repo.GetAllAuthors(ctx)

	if err != nil {
		http.Error(w, "Error getting authors", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func (h *Handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var author queries.CreateAuthorParams

	err := json.NewDecoder(r.Body).Decode(&author)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newAuthor, err := h.Repo.CreateAuthor(ctx, author)

	if err != nil {
		http.Error(w, "Error creating author", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newAuthor)
}
