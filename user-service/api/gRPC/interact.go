package gRPC

import (
	"context"

	"github.com/timoth-y/kicksware-api/service-common/util"

	"github.com/timoth-y/kicksware-api/user-service/api/gRPC/proto"
	"github.com/timoth-y/kicksware-api/user-service/usecase/business"
)

func (h *Handler) Like(ctx context.Context, request *proto.LikeRequest) (resp *proto.LikeResponse, err error) {
	userID, ok := getLikeUserID(ctx, request); if !ok {
		return nil, business.ErrUserNotFound
	}
	err = h.interact.Like(userID, request.EntityID)
	resp = &proto.LikeResponse{
		Success: err == nil,
		Error: util.GetErrorMsg(err),
	}
	return
}

func (h *Handler) Unlike(ctx context.Context, request *proto.LikeRequest) (resp *proto.LikeResponse, err error) {
	userID, ok := getLikeUserID(ctx, request); if !ok {
		return nil, business.ErrUserNotFound
	}
	err = h.interact.Unlike(userID, request.EntityID)

	resp = &proto.LikeResponse{
		Success: err == nil,
		Error: util.GetErrorMsg(err),
	}
	return
}

func getLikeUserID(ctx context.Context, request *proto.LikeRequest) (string, bool) {
	userID := request.UserID
	if id, ok := util.RetrieveUserID(ctx); ok {
		userID = id
	}; if len(userID) == 0 {
		return "", false
	}
	return userID, true
}