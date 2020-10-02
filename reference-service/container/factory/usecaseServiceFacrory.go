package factory

import (
	rest "go.kicksware.com/api/service-common/api/REST"
	"go.kicksware.com/api/service-common/core"

	"go.kicksware.com/api/reference-service/core/repo"
	"go.kicksware.com/api/reference-service/core/service"
	"go.kicksware.com/api/reference-service/env"
	"go.kicksware.com/api/reference-service/usecase/business"
)

func ProvideDataService(repository repo.SneakerReferenceRepository, auth core.AuthService, config env.ServiceConfig) service.SneakerReferenceService {
	return business.NewSneakerReferenceService(repository, auth, config)
}

func ProvideAuthService(config env.ServiceConfig) core.AuthService {
	return rest.NewAuthService(config.Auth)
}