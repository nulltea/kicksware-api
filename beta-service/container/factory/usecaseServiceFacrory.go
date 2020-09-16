package factory

import (
	"github.com/timoth-y/kicksware-api/beta-service/core/repo"
	"github.com/timoth-y/kicksware-api/beta-service/core/service"
	"github.com/timoth-y/kicksware-api/beta-service/env"
	"github.com/timoth-y/kicksware-api/beta-service/usecase/business"
)

func ProvideDataService(repository repo.BetaRepository, config env.ServiceConfig) service.BetaService {
	return business.NewBetaService(repository, config.Common)
}

func ProvideAuthService(config env.ServiceConfig) service.AuthService {
	return business.NewAuthService(config.Auth)
}