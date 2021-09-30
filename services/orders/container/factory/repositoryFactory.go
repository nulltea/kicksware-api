package factory

import (
	"log"

	"go.kicksware.com/api/services/orders/core/repo"
	"go.kicksware.com/api/services/orders/env"
	"go.kicksware.com/api/services/orders/usecase/storage/mongo"
	"go.kicksware.com/api/services/orders/usecase/storage/postgres"
)

func ProvideRepository(config env.ServiceConfig) repo.OrderRepository {
	switch config.Common.UsedDB {
	case "mongo":
		repo, err := mongo.NewRepository(config.Mongo); if err != nil {
			log.Fatal(err)
		}
		return repo
	case "postgres":
		repo, err := postgres.NewRepository(config.Postgres); if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
