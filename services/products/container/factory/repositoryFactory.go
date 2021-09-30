package factory

import (
	"log"

	"go.kicksware.com/api/services/products/core/repo"
	"go.kicksware.com/api/services/products/env"
	"go.kicksware.com/api/services/products/usecase/storage/mongo"
	"go.kicksware.com/api/services/products/usecase/storage/postgres"
	"go.kicksware.com/api/services/products/usecase/storage/redis"
)

func ProvideRepository(config env.ServiceConfig) repo.SneakerProductRepository {
	switch config.Common.UsedDB {
	case "redis":
		repo, err := redis.NewRedisRepository(config.Redis); if err != nil {
			log.Fatal(err)
		}
		return repo
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
