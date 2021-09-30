package gRPC

import (
	"crypto/rsa"

	"go.kicksware.com/api/shared/core/meta"
	"google.golang.org/grpc"

	"go.kicksware.com/api/services/cdn/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterContentServiceServer(server, handler)
	}
}

func (h *Handler) ProvideAccessRoles() meta.AccessConfig {
	return meta.AccessConfig{
		"/proto.ContentService/Original": meta.ZeroAccess,
		"/proto.ContentService/Crop": meta.ZeroAccess,
		"/proto.ContentService/Resize": meta.ZeroAccess,
		"/proto.ContentService/Thumbnail": meta.ZeroAccess,
		"/proto.ContentService/Upload": meta.RegularAccess,
		"/proto.ContentService/StreamContent": meta.ZeroAccess,
	}
}

func (h *Handler) ProvideAuthKey() *rsa.PublicKey {
	return h.auth.PublicKey()
}
