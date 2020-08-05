package startup

import (
	di "github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/container"
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/core"

	conf "beta-service/container/config"
	"beta-service/env"
)

func InitializeServer() (srv core.Server) {
	env.InitEnvironment()
	config, err := env.ReadServiceConfig(env.ServiceConfigPath); if err != nil {
		return nil
	}
	container := di.NewServiceContainer()
	conf.ConfigureContainer(container, config)
	container.Resolve(&srv)
	return
}

