package factory

import (
	"log"

	"reference-service/core/repo"
	"reference-service/env"
	"reference-service/usecase/storage/mongo"
	"reference-service/usecase/storage/postgres"
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
