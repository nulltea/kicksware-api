package rest

import (
	"net/http"
)

func (h *handler) HealthZ(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *handler) ReadyZ(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // TODO DB readiness check
}
