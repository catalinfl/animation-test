package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/catalinfl/blogapi/queries"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) GetCommentsForPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postId := chi.URLParam(r, "postId")
	postIdInt, err := strconv.Atoi(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comments, err := h.Repo.GetCommentsForPost(ctx, int64(postIdInt))
	if err != nil {
		http.Error(w, "Comments error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

func (h *Handler) GetCommentsForAuthor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	authorId := chi.URLParam(r, "authorId")
	authorIdInt, err := strconv.Atoi(authorId)
	if err != nil {
		http.Error(w, "Error at author id", http.StatusInternalServerError)
		return
	}

	comments, err := h.Repo.GetCommentsForAuthor(ctx, int64(authorIdInt))

	if err != nil {
		http.Error(w, "Error at comments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var comment queries.CreateCommentParams
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Error at decoding comment from body", http.StatusInternalServerError)
		return
	}

	commentPost, err := h.Repo.CreateComment(ctx, comment)
	if err != nil {
		http.Error(w, "Error at creating comment", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commentPost)
}
