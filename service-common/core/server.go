package core

import (
	"crypto/rsa"

	"github.com/go-chi/chi"
	"google.golang.org/grpc"

	"go.kicksware.com/api/service-common/core/meta"
)

type Server interface {
	SetupEncryption(cert *meta.TLSCertificate)
	SetupAuth(pb *rsa.PublicKey, accessRoles meta.AccessConfig) // Must be configured before rest & gRPC sub servers!
	SetupREST(router chi.Router) // Setup rest sub server configuration
	SetupGRPC(fn func(srv *grpc.Server)) // Setup gRPC sub server configuration
	SetupLogger()
	Start()
	Shutdown()
}
