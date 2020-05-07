package factory

import (
	"reference-service/core/repo"
	"reference-service/core/service"
	"reference-service/env"
	"reference-service/usecase/business"
)

func ProvideDataService(repository repo.SneakerReferenceRepository, config env.ServiceConfig) service.SneakerReferenceService {
	return business.NewSneakerReferenceService(repository, config.Common)
}

func ProvideAuthService(config env.ServiceConfig) service.AuthService {
	return business.NewAuthService(config.Auth)
}