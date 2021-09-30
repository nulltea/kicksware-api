package factory

import (
	"github.com/go-chi/chi"

	"go.kicksware.com/api/shared/core"
	"go.kicksware.com/api/shared/server"

	"go.kicksware.com/api/services/references/api/gRPC"
	"go.kicksware.com/api/services/references/env"
)

func ProvideServer(config env.ServiceConfig, router chi.Router, handler *gRPC.Handler) core.Server {
	srv := server.NewInstance(config.Common.Host)
	srv.SetupEncryption(config.Security.TLSCertificate)
	srv.SetupAuth(handler.ProvideAuthKey(), handler.ProvideAccessRoles())
	srv.SetupREST(router)
	srv.SetupGRPC(gRPC.ProvideRemoteSetup(handler))
	return srv
}
