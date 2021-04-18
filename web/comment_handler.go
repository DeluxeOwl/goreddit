package web

import (
	"net/http"

	"github.com/DeluxeOwl/goreddit"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type CommentHandler struct {
	store    goreddit.Store
	sessions *scs.SessionManager
}

func (h *CommentHandler) Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := r.FormValue("content")
		idStr := chi.URLParam(r, "postID")

		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.store.CreateComment(&goreddit.Comment{
			ID:      uuid.New(),
			PostID:  id,
			Content: content,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		h.sessions.Put(r.Context(), "flash", "Your comment has been submitted.")

		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}
func (h *CommentHandler) Vote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		c, err := h.store.Comment(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		dir := r.URL.Query().Get("dir")
		if dir == "up" {
			c.Votes++
		} else if dir == "down" {
			c.Votes--
		} else {
			http.Error(w, "Please don't test this parameter :)", http.StatusInternalServerError)
			return
		}

		if err := h.store.UpdateComment(&c); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, r.Referer(), http.StatusFound)
	}
}
