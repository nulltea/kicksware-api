package gRPC

import (
	"crypto/rsa"

	"go.kicksware.com/api/service-common/core/meta"
	"google.golang.org/grpc"

	"go.kicksware.com/api/order-service/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterOrderServiceServer(server, handler)
	}
}

func (h *Handler) ProvideAccessRoles() meta.AccessConfig {
	return meta.AccessConfig{
		"/proto.OrderService/GetOrders": meta.RegularAccess,
		"/proto.OrderService/CountOrders": meta.RegularAccess,
		"/proto.OrderService/AddOrders": meta.UserAccess,
		"/proto.OrderService/EditOrders": meta.AdminAccess,
		"/proto.OrderService/DeleteOrders": meta.AdminAccess,
	}
}

func (h *Handler) ProvideAuthKey() *rsa.PublicKey {
	return h.auth.PublicKey()
}