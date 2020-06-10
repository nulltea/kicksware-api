package factory

import (
	"cdn-service/core/repo"
	"cdn-service/core/service"
	"cdn-service/env"
	"cdn-service/usecase/business"
)

func ProvideContentService(repository repo.ContentRepository, config env.ServiceConfig) service.ContentService {
	return business.NewContentService(repository, config.Common)
}

func ProvideAuthService(config env.ServiceConfig) service.AuthService {
	return business.NewAuthService(config.Auth)
}