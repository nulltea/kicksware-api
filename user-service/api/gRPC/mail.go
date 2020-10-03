package gRPC

import (
	"context"

	common "go.kicksware.com/api/service-common/api/proto"
	"go.kicksware.com/api/service-common/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.kicksware.com/api/user-service/api/gRPC/proto"
	"go.kicksware.com/api/user-service/usecase/business"
)

func (h *Handler) SendEmailConfirmation(ctx context.Context, request *proto.MailRequest) (resp *common.CommonResponse, err error) {
	userID, ok := getMailUserID(ctx, request); if !ok {
		return nil, business.ErrUserNotFound
	}
	err = h.mail.SendEmailConfirmation(userID, request.CallbackURL)
	resp = &common.CommonResponse{
		Success: err == nil,
		Error: util.GetErrorMsg(err),
	}
	return
}

func (h *Handler) SendResetPassword(ctx context.Context, request *proto.MailRequest) (resp *common.CommonResponse, err error) {
	userID, ok := getMailUserID(ctx, request); if !ok {
		return nil, business.ErrUserNotFound
	}
	err = h.mail.SendResetPassword(userID, request.CallbackURL)

	resp = &common.CommonResponse{
		Success: err == nil,
		Error: util.GetErrorMsg(err),
	}
	return
}

func (h *Handler) SendNotification(ctx context.Context, request *proto.MailRequest) (resp *common.CommonResponse, err error) {
	userID, ok := getMailUserID(ctx, request); if !ok {
		return nil, business.ErrUserNotFound
	}
	err = h.mail.SendNotification(userID, request.MessageContent)
	resp = &common.CommonResponse{
		Success: err == nil,
		Error: util.GetErrorMsg(err),
	}
	return
}

func (h *Handler) Subscribe(ctx context.Context, request *proto.SubscribeRequest) (resp *common.CommonResponse, err error) {
	userID, ok := getSubsUserID(ctx, request); if !ok {
		return nil, business.ErrUserNotFound
	}
	if len(request.Email) == 0 {
		return nil, status.Error(codes.InvalidArgument, "email mast be specified in order to subscribe")
	}
	err = h.mail.Subscribe(request.Email, userID);
	resp = &common.CommonResponse{
		Success: err == nil,
		Error: util.GetErrorMsg(err),
	}
	return
}

func (h *Handler) Unsubscribe(ctx context.Context, request *proto.SubscribeRequest) (resp *common.CommonResponse, err error) {
	if len(request.Email) == 0 {
		return nil, status.Error(codes.InvalidArgument, "email mast be specified in order to unsubscribe")
	}
	err = h.mail.Unsubscribe(request.Email);
	resp = &common.CommonResponse{
		Success: err == nil,
		Error: util.GetErrorMsg(err),
	}
	return
}

func getMailUserID(ctx context.Context, request *proto.MailRequest) (string, bool) {
	userID := request.UserID
	if id, ok := util.RetrieveUserID(ctx); ok {
		userID = id
	}; if len(userID) == 0 {
		return "", false
	}
	return userID, true
}

func getSubsUserID(ctx context.Context, request *proto.SubscribeRequest) (string, bool) {
	userID := request.UserID
	if id, ok := util.RetrieveUserID(ctx); ok {
		userID = id
	}; if len(userID) == 0 {
		return "", false
	}
	return userID, true
}