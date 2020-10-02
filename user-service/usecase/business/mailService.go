package business

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net"
	"net/smtp"

	"github.com/golang/glog"
	"github.com/pkg/errors"

	"go.kicksware.com/api/user-service/core/service"
	"go.kicksware.com/api/user-service/env"
)

type mailService struct {
	userService    service.UserService
	config         env.MailConfig
	fallbackConfig env.MailConfig
}

func NewMailService(userService service.UserService, config env.MailConfig, fallbackConfig env.MailConfig) service.MailService {
	return &mailService {
		userService,
		config,
		fallbackConfig,
	}
}

func mailClient(config env.MailConfig) (*smtp.Client, error) {
	host, _, _ := net.SplitHostPort(config.Server)
	tlsConfig := &tls.Config {
		InsecureSkipVerify: true,
		ServerName: config.Server,
	}

	conn, err := tls.Dial("tcp", config.Server, tlsConfig); if err != nil {
		return nil, err
	}

	client, err := smtp.NewClient(conn, host); if err != nil {
		return nil, err
	}

	if err := client.Auth(newEmailAuth(config)); err != nil {
		return nil, err
	}

	return client, nil
}

func newEmailAuth(config env.MailConfig) smtp.Auth {
	host, _, _ := net.SplitHostPort(config.Server)
	return smtp.PlainAuth("", config.Address, config.Password, host)
}

func (s *mailService) SendEmailConfirmation(userID, callbackURL string) error { //
	user, err := s.userService.FetchOne(userID); if err != nil {
		return err
	}
	values := map[string]string{
		"link": callbackURL,
	}
	msg, err := useTemplate(s.config.VerifyEmailTemplate, values); if err != nil {
		return errors.Wrapf(err, "mailService: Could not parse or use specified template %q", s.config.VerifyEmailTemplate)
	}
	return s.sendMail("Kicksware account verification", msg, user.Email)
}

func (s *mailService) SendResetPassword(userID, callbackURL string) error {
	user, err := s.userService.FetchOne(userID); if err != nil {
		return err
	}
	values := map[string]string{
		"link": callbackURL,
	}
	msg, err := useTemplate(s.config.ResetPasswordTemplate, values); if err != nil {
		return errors.Wrapf(err, "mailService: Could not parse or use specified template %q", s.config.ResetPasswordTemplate)
	}
	return s.sendMail("Kicksware password reset", msg, user.Email)
}

func (s *mailService) SendNotification(userID, notificationContent string) error {
	panic("implement me")
}

func (s *mailService) sendMail(subject string, msg string, to string) error {
	if err := trySendMail(s.config, subject, msg, to); err != nil {
		glog.Errorln(err)
	} else {
		return nil
	}

	if err := trySendMail(s.fallbackConfig, subject, msg, to);  err != nil {
		glog.Errorln(err)
		return err
	}

	return nil
}

func trySendMail(config env.MailConfig, subject string, msg string, to string) error {
	client, err := mailClient(config); if err != nil {
		return err
	}

	if err := client.Mail(config.Address); err != nil {
		return err
	}

	if err := client.Rcpt(to); err != nil {
		return err
	}

	w, err := client.Data(); if err != nil {
		return err
	}

	body := formEmailRequestBody(subject, msg, config.Address, to)
	_, err = w.Write(body); if err != nil {
		return err
	}

	err = w.Close(); if err != nil {
		return err
	}
	return nil
}

func useTemplate(path string, format interface{}) (string, error) {
	var w bytes.Buffer
	tmpl, err := template.ParseFiles(path); if err != nil {
		return "", err
	}
	if err = tmpl.Execute(&w, format); err != nil {
		return "", err
	}
	return w.String(), nil
}

func formEmailRequestBody(subject, body string, from, to string) []byte {
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = subject
	headers["MIME-version"] = "1.0"
	headers["Content-Type"] = "text/html"
	headers["charset"] = "\"UTF-8\""

	message := ""
	for key := range headers {
		message += fmt.Sprintf("%s: %s\r\n", key, headers[key])
	}
	message += "\r\n" + body

	return []byte(message)
}