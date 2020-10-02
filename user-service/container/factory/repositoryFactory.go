package factory

import (
	"log"

	"go.kicksware.com/api/user-service/core/repo"
	"go.kicksware.com/api/user-service/env"
	"go.kicksware.com/api/user-service/usecase/storage/mongo"
	"go.kicksware.com/api/user-service/usecase/storage/postgres"
	"go.kicksware.com/api/user-service/usecase/storage/redis"
)

func ProvideRepository(config env.ServiceConfig) repo.UserRepository {
	switch config.Common.UsedDB {
	case "redis":
		repo, err := redis.NewRepository(config.Redis); if err != nil {
			log.Fatal(err)
		}
		return repo
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

func ProvideLikesRepository(config env.ServiceConfig) repo.LikesRepository {
	switch config.Common.UsedDB {
	case "mongo":
		repo, err := mongo.NewLikesRepository(config.Mongo); if err != nil {
			log.Fatal(err)
		}
		return repo
	case "postgres":
		repo, err := postgres.NewLikesRepository(config.Postgres); if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}

func ProvideRemotesRepository(config env.ServiceConfig) repo.RemoteRepository {
	switch config.Common.UsedDB {
	case "mongo":
		repo, err := mongo.NewRemoteRepository(config.Mongo); if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}