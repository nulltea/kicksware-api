package factory

import (
	"user-service/core/repo"
	"user-service/core/service"
	"user-service/env"
	"user-service/usecase/business"
)

func ProvideDataService(repository repo.UserRepository, remoteRepository repo.RemoteRepository) service.UserService {
	return business.NewUserService(repository, remoteRepository)
}

func ProvideAuthService(service service.UserService, config env.ServiceConfig) service.AuthService {
	return business.NewAuthServiceJWT(
		service,
		config.Auth,
	)
}

func ProvideMailService(service service.UserService, config env.ServiceConfig) service.MailService {
	return business.NewMailService(
		service,
		config.Mail,
		config.FallbackMail,
	)
}

func ProvideInteractService(service service.UserService, likesRepo repo.LikesRepository) service.InteractService {
	return business.NewInteractService(
		service,
		likesRepo,
	)
}