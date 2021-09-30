package factory

import (
	"go.kicksware.com/api/shared/core"
	"go.kicksware.com/api/shared/server"

	"go.kicksware.com/api/services/rating/env"
)

func ProvideServer(config env.ServiceConfig) core.Server {
	srv := server.NewInstance(config.Common.Host)
	return srv
}
