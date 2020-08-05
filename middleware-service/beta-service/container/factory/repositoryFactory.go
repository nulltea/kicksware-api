package factory

import (
	"log"

	"beta-service/core/repo"
	"beta-service/env"
	"beta-service/usecase/storage/mongo"
	"beta-service/usecase/storage/postgres"
)

func ProvideRepository(config env.ServiceConfig) repo.BetaRepository {
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
