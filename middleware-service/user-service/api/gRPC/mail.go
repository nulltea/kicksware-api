package gRPC

import (
	"context"

	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/api/gRPC/proto"
)

func (h *Handler) SendEmailConfirmation(ctx context.Context, request *proto.MailRequest) (resp *proto.MailResponse, err error) {
	err = h.mail.SendEmailConfirmation(request.UserID, request.CallbackURL)
	resp = &proto.MailResponse{
		Success: err == nil,
		Error: err.Error(),
	}
	return
}

func (h *Handler) SendResetPassword(ctx context.Context, request *proto.MailRequest) (resp *proto.MailResponse, err error) {
	err = h.mail.SendResetPassword(request.UserID, request.CallbackURL)
	resp = &proto.MailResponse{
		Success: err == nil,
		Error: err.Error(),
	}
	return
}

func (h *Handler) SendNotification(ctx context.Context, request *proto.MailRequest) (resp *proto.MailResponse, err error) {
	err = h.mail.SendNotification(request.UserID, request.MessageContent)
	resp = &proto.MailResponse{
		Success: err == nil,
		Error: err.Error(),
	}
	return
}
