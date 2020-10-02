package factory

import (
	"log"

	"go.kicksware.com/api/product-service/core/repo"
	"go.kicksware.com/api/product-service/env"
	"go.kicksware.com/api/product-service/usecase/storage/mongo"
	"go.kicksware.com/api/product-service/usecase/storage/postgres"
	"go.kicksware.com/api/product-service/usecase/storage/redis"
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
