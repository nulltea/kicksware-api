package core

import (
	"crypto/rsa"

	"github.com/go-chi/chi"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/model"
	"google.golang.org/grpc"
)

type Server interface {
	SetupAuth(pb *rsa.PublicKey, accessRoles map[string][]model.UserRole) // Must be configured before REST & gRPC sub servers!
	SetupREST(router chi.Router) // Setup REST sub server configuration
	SetupGRPC(fn func(srv *grpc.Server)) // Setup gRPC sub server configuration
	Start()
	Shutdown()
}
