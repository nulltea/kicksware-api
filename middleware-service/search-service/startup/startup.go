package startup

import (
	"log"

	"github.com/pkg/errors"
	di "github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/container"
	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/core"

	conf "search-service/container/config"
	"search-service/core/service"
	"search-service/env"
)

func InitializeServer() (srv core.Server, container di.ServiceContainer) {
	env.InitEnvironment()
	config, err := env.ReadServiceConfig(env.ServiceConfigPath); if err != nil {
		return
	}
	container = di.NewServiceContainer()
	conf.ConfigureContainer(container, config)
	container.Resolve(&srv)
	return
}

func PerformDataSync(container di.ServiceContainer) error {
	return container.ResolveFor(func (service service.ReferenceSyncService) error {
		if err := service.SyncAll(nil); err != nil {
			log.Println(errors.Wrap(err, "search-service::startup.PerformDataSync: sneaker references replication sync failed"))
			return errors.Wrap(err, "search-service::startup.PerformDataSync: sneaker references replication sync failed")
		}
		log.Println("search-service::startup.PerformDataSync: sneaker references replication sync completed with success")
		return nil
	})
}