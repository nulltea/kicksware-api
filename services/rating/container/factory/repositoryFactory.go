package factory

import (
	"log"

	"go.kicksware.com/api/services/rating/core/repo"
	"go.kicksware.com/api/services/rating/env"
	"go.kicksware.com/api/services/rating/usecase/storage/mongo"
)

func ProvideRepository(config env.ServiceConfig) repo.RatingRepository {
	repo, err := mongo.NewRepository(config.Mongo); if err != nil {
		log.Fatal(err)
	}
	return repo
}
