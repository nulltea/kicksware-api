package rest

import (
	"net/http"
)

func (h *Handler) HealthZ(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ReadyZ(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // TODO DB readiness check
}
