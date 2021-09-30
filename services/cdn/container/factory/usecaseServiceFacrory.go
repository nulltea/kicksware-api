package factory

import (
	"go.kicksware.com/api/shared/api/rest"
	"go.kicksware.com/api/shared/core"

	"go.kicksware.com/api/services/cdn/core/repo"
	"go.kicksware.com/api/services/cdn/core/service"
	"go.kicksware.com/api/services/cdn/env"
	"go.kicksware.com/api/services/cdn/usecase/business"
)

func ProvideContentService(repository repo.ContentRepository, config env.ServiceConfig) service.ContentService {
	return business.NewContentService(repository, config.Common)
}

func ProvideAuthService(config env.ServiceConfig) core.AuthService {
	return rest.NewAuthService(config.Auth)
}
