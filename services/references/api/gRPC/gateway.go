package gRPC

import (
	"crypto/rsa"

	"go.kicksware.com/api/shared/core/meta"
	"google.golang.org/grpc"

	"go.kicksware.com/api/services/references/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterReferenceServiceServer(server, handler)
	}
}

func (h *Handler) ProvideAccessRoles() meta.AccessConfig {
	return meta.AccessConfig{
		"/proto.ReferenceService/GetReferences": meta.RegularAccess,
		"/proto.ReferenceService/CountReferences": meta.RegularAccess,
		"/proto.ReferenceService/AddReferences": meta.AdminAccess,
		"/proto.ReferenceService/EditReferences": meta.AdminAccess,
		"/proto.ReferenceService/DeleteReferences": meta.AdminAccess,
	}
}

func (h *Handler) ProvideAuthKey() *rsa.PublicKey {
	return h.auth.PublicKey()
}
