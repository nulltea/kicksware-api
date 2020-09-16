package factory

import (
	"log"

	"github.com/timoth-y/kicksware-api/reference-service/core/repo"
	"github.com/timoth-y/kicksware-api/reference-service/env"
	"github.com/timoth-y/kicksware-api/reference-service/usecase/storage/mongo"
	"github.com/timoth-y/kicksware-api/reference-service/usecase/storage/postgres"
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
