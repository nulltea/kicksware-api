package rest

import (
	"net/http"
)

func (h *Handler) SendEmailConfirmation(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userID := query.Get("userID")
	callbackURL := query.Get("callbackURL")
	err := h.mail.SendEmailConfirmation(userID, callbackURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *Handler) SendResetPassword(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userID := query.Get("userID")
	callbackURL := query.Get("callbackURL")
	err := h.mail.SendResetPassword(userID, callbackURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *Handler) SendNotification(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userID := query.Get("userID")
	callbackURL := query.Get("content")
	err := h.mail.SendNotification(userID, callbackURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}