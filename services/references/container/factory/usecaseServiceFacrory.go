package factory

import (
	rest "go.kicksware.com/api/shared/api/rest"
	"go.kicksware.com/api/shared/core"

	"go.kicksware.com/api/services/references/core/repo"
	"go.kicksware.com/api/services/references/core/service"
	"go.kicksware.com/api/services/references/env"
	"go.kicksware.com/api/services/references/usecase/business"
)

func ProvideDataService(repository repo.SneakerReferenceRepository, auth core.AuthService, config env.ServiceConfig) service.SneakerReferenceService {
	return business.NewSneakerReferenceService(repository, auth, config)
}

func ProvideAuthService(config env.ServiceConfig) core.AuthService {
	return rest.NewAuthService(config.Auth)
}
