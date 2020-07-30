package rest

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) Like(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.User.Username()
	entityID := chi.URLParam(r, "entityID")
	err := h.interact.Like(userID, entityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *Handler) Unlike(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.User.Username()
	entityID := chi.URLParam(r, "entityID")
	err := h.interact.Unlike(userID, entityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}
