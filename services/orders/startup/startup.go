package startup

import (
	di "go.kicksware.com/api/shared/container"
	"go.kicksware.com/api/shared/core"

	conf "go.kicksware.com/api/services/orders/container/config"
	"go.kicksware.com/api/services/orders/env"
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
