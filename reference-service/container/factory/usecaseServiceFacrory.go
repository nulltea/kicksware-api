package factory

import (
	"github.com/timoth-y/kicksware-api/reference-service/core/repo"
	"github.com/timoth-y/kicksware-api/reference-service/core/service"
	"github.com/timoth-y/kicksware-api/reference-service/env"
	"github.com/timoth-y/kicksware-api/reference-service/usecase/business"
)

func ProvideDataService(repository repo.SneakerReferenceRepository, config env.ServiceConfig) service.SneakerReferenceService {
	return business.NewSneakerReferenceService(repository, config.Common)
}

func ProvideAuthService(config env.ServiceConfig) service.AuthService {
	return business.NewAuthService(config.Auth)
}