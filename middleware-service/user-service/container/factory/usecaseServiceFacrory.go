package factory

import (
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/repo"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/service"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/env"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/usecase/business"
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