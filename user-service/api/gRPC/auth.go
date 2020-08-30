package gRPC

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/api/gRPC/proto"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/meta"
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

func (h *Handler) Guest(ctx context.Context, empty *empty.Empty) (resp *proto.AuthToken, err error) {
	var token *meta.AuthToken;

	token, err = h.auth.Guest(); if err != nil || token == nil {
		return
	}

	resp = proto.AuthToken{}.FromNative(token)
	return
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