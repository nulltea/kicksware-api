package core

import (
	"github.com/go-chi/chi"
)

type Server interface {
	SetupRoutes(router chi.Router)
	Start()
	Shutdown()
}
