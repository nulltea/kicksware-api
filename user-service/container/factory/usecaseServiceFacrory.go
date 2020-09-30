package factory

import (
	"github.com/timoth-y/kicksware-api/user-service/core/repo"
	"github.com/timoth-y/kicksware-api/user-service/core/service"
	"github.com/timoth-y/kicksware-api/user-service/env"
	"github.com/timoth-y/kicksware-api/user-service/usecase/business"
)

func ProvideDataService(repository repo.UserRepository, remoteRepository repo.RemoteRepository, config env.ServiceConfig) service.UserService {
	return business.NewUserService(repository, remoteRepository, config)
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