package service

type MailService interface {
	SendEmailConfirmation(userID, callbackURL string)
	SendResetPassword(userID, callbackURL string)
	SendNotification(userID, notificationContent string)
}