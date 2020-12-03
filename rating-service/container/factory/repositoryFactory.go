package factory

import (
	"log"

	"go.kicksware.com/api/rating-service/core/repo"
	"go.kicksware.com/api/rating-service/env"
	"go.kicksware.com/api/rating-service/usecase/storage/mongo"
)

func ProvideRepository(config env.ServiceConfig) repo.RatingRepository {
	repo, err := mongo.NewRepository(config.Mongo); if err != nil {
		log.Fatal(err)
	}
	return repo
}
