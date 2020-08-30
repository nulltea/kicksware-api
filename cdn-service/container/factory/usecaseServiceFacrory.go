package factory

import (
	"github.com/timoth-y/kicksware-platform/middleware-service/cdn-service/core/repo"
	"github.com/timoth-y/kicksware-platform/middleware-service/cdn-service/core/service"
	"github.com/timoth-y/kicksware-platform/middleware-service/cdn-service/env"
	"github.com/timoth-y/kicksware-platform/middleware-service/cdn-service/usecase/business"
)

func ProvideContentService(repository repo.ContentRepository, config env.ServiceConfig) service.ContentService {
	return business.NewContentService(repository, config.Common)
}

func ProvideAuthService(config env.ServiceConfig) service.AuthService {
	return business.NewAuthService(config.Auth)
}