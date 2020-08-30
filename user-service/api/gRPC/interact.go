package gRPC

import (
	"context"

	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/api/gRPC/proto"
)

func (h *Handler) Like(ctx context.Context, request *proto.LikeRequest) (resp *proto.LikeResponse, err error) {
	err = h.interact.Like(request.UserID, request.EntityID)
	resp = &proto.LikeResponse{
		Success: err == nil,
		Error: err.Error(),
	}
	return
}

func (h *Handler) Unlike(ctx context.Context, request *proto.LikeRequest) (resp *proto.LikeResponse, err error) {
	err = h.interact.Unlike(request.UserID, request.EntityID)
	resp = &proto.LikeResponse{
		Success: err == nil,
		Error: err.Error(),
	}
	return
}