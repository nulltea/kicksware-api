package factory

import (
	"github.com/go-chi/chi"
	"go.kicksware.com/api/service-common/core"
	"go.kicksware.com/api/service-common/server"

	grpc "go.kicksware.com/api/user-service/api/gRPC"
	"go.kicksware.com/api/user-service/env"
)

func ProvideServer(config env.ServiceConfig, router chi.Router, handler *grpc.Handler) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupEncryption(config.Security.TLSCertificate)
	// srv.SetupAuth(handler.ProvideAuthKey(), handler.ProvideAccessRoles())
	srv.SetupREST(router)
	srv.SetupGRPC(grpc.ProvideRemoteSetup(handler))
	return srv
}