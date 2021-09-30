package gRPC

import (
	"crypto/rsa"

	"go.kicksware.com/api/shared/core/meta"
	"google.golang.org/grpc"

	"go.kicksware.com/api/services/products/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterProductServiceServer(server, handler)
	}
}

func (h *Handler) ProvideAccessRoles() meta.AccessConfig {
	return meta.AccessConfig{
		"/proto.ProductService/GetProducts": meta.RegularAccess,
		"/proto.ProductService/CountProducts": meta.RegularAccess,
		"/proto.ProductService/AddProducts": meta.UserAccess,
		"/proto.ProductService/EditProducts": meta.UserAccess,
		"/proto.ProductService/DeleteProducts": meta.UserAccess,
	}
}

func (h *Handler) ProvideAuthKey() *rsa.PublicKey {
	return h.auth.PublicKey()
}
