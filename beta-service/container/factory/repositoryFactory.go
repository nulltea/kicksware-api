package factory

import (
	"log"

	"go.kicksware.com/api/beta/core/repo"
	"go.kicksware.com/api/beta/env"
	"go.kicksware.com/api/beta/usecase/storage/mongo"
	"go.kicksware.com/api/beta/usecase/storage/postgres"
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
