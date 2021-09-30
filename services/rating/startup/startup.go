package startup

import (
	di "go.kicksware.com/api/shared/container"

	"go.kicksware.com/api/services/rating/api/events"
	conf "go.kicksware.com/api/services/rating/container/config"
	"go.kicksware.com/api/services/rating/env"
)

func InitializeEventBus() (bus *events.Handler) {
	env.InitEnvironment()
	config, err := env.ReadServiceConfig(env.ServiceConfigPath); if err != nil {
		return nil
	}
	container := di.NewServiceContainer()
	conf.ConfigureContainer(container, config)
	container.Resolve(&bus)
	return
}
