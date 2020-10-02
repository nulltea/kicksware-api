package core

import (
	"crypto/rsa"

	"github.com/go-chi/chi"
	"go.kicksware.com/api/user-service/core/model"
	"google.golang.org/grpc"

	"go.kicksware.com/api/service-common/core/meta"
)

type Server interface {
	SetupEncryption(cert *meta.TLSCertificate)
	SetupAuth(pb *rsa.PublicKey, accessRoles map[string][]model.UserRole) // Must be configured before rest & gRPC sub servers!
	SetupREST(router chi.Router) // Setup rest sub server configuration
	SetupGRPC(fn func(srv *grpc.Server)) // Setup gRPC sub server configuration
	SetupLogger()
	Start()
	Shutdown()
}
