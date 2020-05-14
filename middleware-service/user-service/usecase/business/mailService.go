package business

import (
	"net/smtp"

	"user-service/core/service"
	"user-service/env"
)

type mailService struct {
	userService service.UserService
	config env.MailConfig
	client *smtp.Client
}

func NewMailService(userService service.UserService, config env.MailConfig) service.MailService {
	client, err := newEmailClient(config); if err != nil {
		return nil
	}
	return &mailService {
		userService,
		config,
		client,
	}
}

func newEmailClient(config env.MailConfig) (*smtp.Client, error) {
	return nil, nil
}

func (m mailService) SendEmailConfirmation(userID, callbackURL string) {
	panic("implement me")
}

func (m mailService) SendResetPassword(userID, callbackURL string) {
	panic("implement me")
}

func (m mailService) SendNotification(userID, notificationContent string) {
	panic("implement me")
}