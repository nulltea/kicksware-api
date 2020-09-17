package core

import (
	"crypto/rsa"

	"github.com/go-chi/chi"
	// "github.com/timoth-y/kicksware-api/user-service/core/model"
	"google.golang.org/grpc"

	"github.com/timoth-y/kicksware-api/service-common/core/meta"
)

type Server interface {
	SetupEncryption(cert *meta.TLSCertificate)
	SetupAuth(pb *rsa.PublicKey, accessRoles map[string][]model.UserRole) // Must be configured before REST & gRPC sub servers!
	SetupREST(router chi.Router) // Setup REST sub server configuration
	SetupGRPC(fn func(srv *grpc.Server)) // Setup gRPC sub server configuration
	Start()
	Shutdown()
}
