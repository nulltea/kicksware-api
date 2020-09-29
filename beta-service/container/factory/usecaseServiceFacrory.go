package factory

import (
	"go.kicksware.com/kicksware/api/beta-service/core/repo"
	"go.kicksware.com/kicksware/api/beta-service/core/service"
	"go.kicksware.com/kicksware/api/beta-service/env"
	"go.kicksware.com/kicksware/api/beta-service/usecase/business"
)

func ProvideDataService(repository repo.BetaRepository, config env.ServiceConfig) service.BetaService {
	return business.NewBetaService(repository, config.Common)
}

func ProvideAuthService(config env.ServiceConfig) service.AuthService {
	return business.NewAuthService(config.Auth)
}