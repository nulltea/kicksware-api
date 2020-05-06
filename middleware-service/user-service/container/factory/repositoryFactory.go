package factory

import (
	"log"

	"user-service/core/repo"
	"user-service/env"
	"user-service/usecase/storage/mongo"
	"user-service/usecase/storage/postgres"
	"user-service/usecase/storage/redis"
)

func ProvideRepository(config env.ServiceConfig) repo.UserRepository {
	switch config.Common.UsedDB {
	case "redis":
		repo, err := redis.NewRedisRepository(config.Redis.URL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "mongo":
		repo, err := mongo.NewMongoRepository(
			config.Mongo.URL,
			config.Mongo.Database,
			config.Mongo.Collection,
			config.Mongo.Timeout,
		); if err != nil {
			log.Fatal(err)
		}
		return repo
	case "postgres":
		repo, err := postgres.NewPostgresRepository(
			config.Postgres.URL,
			config.Postgres.Collection,
		); if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
