package factory

import (
	"go.kicksware.com/api/service-common/api/rest"
	"go.kicksware.com/api/service-common/core"

	"go.kicksware.com/api/cdn-service/core/repo"
	"go.kicksware.com/api/cdn-service/core/service"
	"go.kicksware.com/api/cdn-service/env"
	"go.kicksware.com/api/cdn-service/usecase/business"
)

func ProvideContentService(repository repo.ContentRepository, config env.ServiceConfig) service.ContentService {
	return business.NewContentService(repository, config.Common)
}

func ProvideAuthService(config env.ServiceConfig) core.AuthService {
	return rest.NewAuthService(config.Auth)
}