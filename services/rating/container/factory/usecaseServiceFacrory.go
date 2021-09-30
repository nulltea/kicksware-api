package factory

import (
	"go.kicksware.com/api/shared/api/rest"
	"go.kicksware.com/api/shared/core"

	"go.kicksware.com/api/services/rating/core/repo"
	"go.kicksware.com/api/services/rating/core/service"
	"go.kicksware.com/api/services/rating/env"
	"go.kicksware.com/api/services/rating/usecase/business"
)

func ProvideDataService(repository repo.RatingRepository, auth core.AuthService, config env.ServiceConfig) service.RatingService {
	return business.NewRatingService(repository, auth, config)
}

func ProvideAuthService(config env.ServiceConfig) core.AuthService {
	return rest.NewAuthService(config.Auth)
}
