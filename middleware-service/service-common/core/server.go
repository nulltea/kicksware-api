package core

import (
	"crypto/rsa"

	"github.com/go-chi/chi"
	"google.golang.org/grpc"
)

type Server interface {
	SetupAuth(pb *rsa.PublicKey, accessRoles map[string][]string) // Must be configured before REST & gRPC sub servers!
	SetupREST(router chi.Router) // Setup REST sub server configuration
	SetupGRPC(fn func(srv *grpc.Server)) // Setup gRPC sub server configuration
	Start()
	Shutdown()
}
