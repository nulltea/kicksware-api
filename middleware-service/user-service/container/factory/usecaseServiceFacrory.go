package factory

import (
	"user-service/core/repo"
	"user-service/core/service"
	"user-service/env"
	"user-service/usecase/business"
)

func ProvideDataService(repository repo.UserRepository) service.UserService {
	return business.NewUserService(repository)
}

func ProvideAuthService(service service.UserService, config env.ServiceConfig) service.AuthService {
	return business.NewAuthServiceJWT(
		service,
		config.Auth.TokenExpirationDelta,
		config.Auth.PrivateKeyPath,
		config.Auth.PublicKeyPath,
	)
}