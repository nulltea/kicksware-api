package factory

import (
	"log"

	"go.kicksware.com/api/services/references/core/repo"
	"go.kicksware.com/api/services/references/env"
	"go.kicksware.com/api/services/references/usecase/storage/mongo"
	"go.kicksware.com/api/services/references/usecase/storage/postgres"
)

func ProvideRepository(config env.ServiceConfig) repo.SneakerReferenceRepository {
	switch config.Common.UsedDB {
	case "mongo":
		repo, err := mongo.NewMongoRepository(config.Mongo); if err != nil {
			log.Fatal(err)
		}
		return repo
	case "postgres":
		repo, err := postgres.NewPostgresRepository(config.Postgres); if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
