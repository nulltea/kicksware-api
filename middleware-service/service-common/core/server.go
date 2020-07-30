package core

import (
	"github.com/go-chi/chi"
	"google.golang.org/grpc"
)

type Server interface {
	SetupREST(router chi.Router)
	SetupRoutes(router chi.Router) // Deprecated: SetupRoutes is deprecated. Use SetupREST instead
	SetupGRPC(fn func(srv *grpc.Server))
	Start()
	Shutdown()
}
