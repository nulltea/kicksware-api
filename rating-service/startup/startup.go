package startup

import (
	di "go.kicksware.com/api/service-common/container"

	"go.kicksware.com/api/rating-service/api/events"
	conf "go.kicksware.com/api/rating-service/container/config"
	"go.kicksware.com/api/rating-service/env"
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
