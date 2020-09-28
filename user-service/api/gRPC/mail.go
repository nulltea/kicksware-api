package gRPC

import (
	"context"

	"github.com/timoth-y/kicksware-api/service-common/service/gRPC"

	"github.com/timoth-y/kicksware-api/user-service/api/gRPC/proto"
	"github.com/timoth-y/kicksware-api/user-service/usecase/business"
)

func (h *Handler) SendEmailConfirmation(ctx context.Context, request *proto.MailRequest) (resp *proto.MailResponse, err error) {
	userID, ok := getMailUserID(ctx, request); if !ok {
		return nil, business.ErrUserNotFound
	}
	err = h.mail.SendEmailConfirmation(userID, request.CallbackURL)
	resp = &proto.MailResponse{
		Success: err == nil,
		Error: getErrorMsg(err),
	}
	return
}

func (h *Handler) SendResetPassword(ctx context.Context, request *proto.MailRequest) (resp *proto.MailResponse, err error) {
	userID, ok := getMailUserID(ctx, request); if !ok {
		return nil, business.ErrUserNotFound
	}
	err = h.mail.SendResetPassword(userID, request.CallbackURL)
	resp = &proto.MailResponse{
		Success: err == nil,
		Error: getErrorMsg(err),
	}
	return
}

func (h *Handler) SendNotification(ctx context.Context, request *proto.MailRequest) (resp *proto.MailResponse, err error) {
	userID, ok := getMailUserID(ctx, request); if !ok {
		return nil, business.ErrUserNotFound
	}
	err = h.mail.SendNotification(userID, request.MessageContent)
	resp = &proto.MailResponse{
		Success: err == nil,
		Error: getErrorMsg(err),
	}
	return
}

func getErrorMsg(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func getMailUserID(ctx context.Context, request *proto.MailRequest) (string, bool) {
	userID := request.UserID
	if id, ok := gRPC.RetrieveUserID(ctx); ok {
		userID = id
	}; if len(userID) == 0 {
		return "", false
	}
	return userID, true
}