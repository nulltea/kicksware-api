package factory

import (
	rest "github.com/timoth-y/kicksware-api/service-common/api/REST"
	"github.com/timoth-y/kicksware-api/service-common/core"

	"github.com/timoth-y/kicksware-api/reference-service/core/repo"
	"github.com/timoth-y/kicksware-api/reference-service/core/service"
	"github.com/timoth-y/kicksware-api/reference-service/env"
	"github.com/timoth-y/kicksware-api/reference-service/usecase/business"
)

func ProvideDataService(repository repo.SneakerReferenceRepository, auth core.AuthService, config env.ServiceConfig) service.SneakerReferenceService {
	return business.NewSneakerReferenceService(repository, auth, config)
}

func ProvideAuthService(config env.ServiceConfig) core.AuthService {
	return rest.NewAuthService(config.Auth)
}