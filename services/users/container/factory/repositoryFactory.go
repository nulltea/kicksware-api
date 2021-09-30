package factory

import (
	"log"

	"go.kicksware.com/api/services/users/core/repo"
	"go.kicksware.com/api/services/users/env"
	"go.kicksware.com/api/services/users/usecase/storage/mongo"
	"go.kicksware.com/api/services/users/usecase/storage/postgres"
	"go.kicksware.com/api/services/users/usecase/storage/redis"
)

func ProvideRepository(config env.ServiceConfig) repo.UserRepository {
	switch config.Common.UsedDB {
	case "redis":
		repo, err := redis.NewRepository(config.UsersDB); if err != nil {
			log.Fatal(err)
		}
		return repo
	case "mongo":
		repo, err := mongo.NewRepository(config.UsersDB); if err != nil {
			log.Fatal(err)
		}
		return repo
	case "postgres":
		repo, err := postgres.NewRepository(config.UsersDB); if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}

func ProvideLikesRepository(config env.ServiceConfig) repo.LikesRepository {
	switch config.Common.UsedDB {
	case "mongo":
		repo, err := mongo.NewLikesRepository(config.LikesDB); if err != nil {
			log.Fatal(err)
		}
		return repo
	case "postgres":
		repo, err := postgres.NewLikesRepository(config.LikesDB); if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}

func ProvideRemotesRepository(config env.ServiceConfig) repo.RemoteRepository {
	switch config.Common.UsedDB {
	case "mongo":
		repo, err := mongo.NewRemoteRepository(config.RemotesDB); if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}

func ProvideSubscriptionsRepository(config env.ServiceConfig) repo.SubscriptionRepository {
	switch config.Common.UsedDB {
	case "mongo":
		repo, err := mongo.NewSubscriptionsRepository(config.SubscriptionsDB); if err != nil {
		log.Fatal(err)
	}
		return repo
	}
	return nil
}
