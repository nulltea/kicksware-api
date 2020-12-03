package factory

import (
	"go.kicksware.com/api/service-common/api/rest"
	"go.kicksware.com/api/service-common/core"

	"go.kicksware.com/api/rating-service/core/repo"
	"go.kicksware.com/api/rating-service/core/service"
	"go.kicksware.com/api/rating-service/env"
	"go.kicksware.com/api/rating-service/usecase/business"
)

func ProvideDataService(repository repo.RatingRepository, auth core.AuthService, config env.ServiceConfig) service.RatingService {
	return business.NewRatingService(repository, auth, config)
}

func ProvideAuthService(config env.ServiceConfig) core.AuthService {
	return rest.NewAuthService(config.Auth)
}