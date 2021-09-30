package gRPC

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.kicksware.com/api/services/users/api/gRPC/proto"
	"go.kicksware.com/api/services/users/core/meta"
	"go.kicksware.com/api/services/users/usecase/business"
)

func (h *Handler) SignUp(ctx context.Context, user *proto.User) (resp *proto.AuthToken, err error) {
	var token *meta.AuthToken; if user == nil {
		return
	}

	token, err = h.auth.SingUp(user.ToNative()); if err != nil || token == nil {
		return
	}

	resp = proto.AuthToken{}.FromNative(token)
	return
}

func (h *Handler) Login(ctx context.Context, user *proto.User) (resp *proto.AuthToken, err error) {
	var token *meta.AuthToken; if user == nil {
		return
	}

	token, err = h.auth.Login(user.ToNative()); if err != nil || token == nil {
		return
	}

	resp = proto.AuthToken{}.FromNative(token)
	return
}

func (h *Handler) Remote(ctx context.Context, user *proto.User) (resp *proto.AuthToken, err error) {
	var token *meta.AuthToken; if user == nil {
		return
	}

	token, err = h.auth.Remote(user.ToNative()); if err != nil || token == nil {
		return
	}

	resp = proto.AuthToken{}.FromNative(token)
	return
}

func (h *Handler) Guest(ctx context.Context, access *proto.AccessKey) (*proto.AuthToken, error) {
	if !h.auth.VerifyAccessKey(access.Key) {
		return nil,status.Errorf(codes.Internal,
			"Error occurred while generating auth token: %v", business.ErrInvalidAccessKey.Error(),
		)
	}

	token, err := h.auth.Guest(); if err != nil || token == nil {
		return nil, status.Errorf(codes.Internal,
			"Error occurred while generating auth token: %v", err,
		)
	}

	return proto.AuthToken{}.FromNative(token), nil
}

func (h *Handler) GenerateToken(ctx context.Context, user *proto.User) (resp *proto.AuthToken, err error) {
	var token *meta.AuthToken; if user == nil {
		return
	}

	token, err = h.auth.GenerateToken(user.ToNative()); if err != nil || token == nil {
		return
	}

	resp = proto.AuthToken{}.FromNative(token)
	return
}

func (h *Handler) Refresh(ctx context.Context, token *proto.AuthToken) (resp *proto.AuthToken, err error) {
	var native *meta.AuthToken; if token == nil {
		return
	}

	native, err = h.auth.Refresh(token.ToNative().Token); if err != nil || token == nil {
		return
	}

	resp = proto.AuthToken{}.FromNative(native)
	return
}

func (h *Handler) Logout(ctx context.Context, token *proto.AuthToken) (resp *proto.AuthToken, err error) {
	err = h.auth.Logout(token.ToNative().Token); if err != nil || token == nil {
		return
	}

	resp = &proto.AuthToken{
		Success: err == nil,
	}
	return
}
