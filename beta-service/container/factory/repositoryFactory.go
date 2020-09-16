package factory

import (
	"log"

	"github.com/timoth-y/kicksware-api/beta-service/core/repo"
	"github.com/timoth-y/kicksware-api/beta-service/env"
	"github.com/timoth-y/kicksware-api/beta-service/usecase/storage/mongo"
	"github.com/timoth-y/kicksware-api/beta-service/usecase/storage/postgres"
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
