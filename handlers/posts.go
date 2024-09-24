package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/catalinfl/blogapi/queries"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := h.Repo.GetPosts(ctx)
	if err != nil {
		http.Error(w, "Error getting posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	post, err := h.Repo.GetPost(ctx, int64(idint))
	if err != nil {
		http.Error(w, "Error getting post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var post queries.CreatePostParams
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid post", http.StatusBadRequest)
		return
	}
	newPost, err := h.Repo.CreatePost(ctx, post)
	if err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPost)
}
