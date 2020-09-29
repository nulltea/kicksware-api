package gRPC

import (
	"crypto/rsa"

	"github.com/timoth-y/kicksware-api/user-service/core/model"
	"google.golang.org/grpc"

	"github.com/timoth-y/kicksware-api/order-service/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterOrderServiceServer(server, handler)
	}
}

func (h *Handler) ProvideAccessRoles() map[string][]model.UserRole {
	roleMap := make(map[string][]model.UserRole)

	regularAccess := []model.UserRole{ model.Guest, model.Regular, model.Admin }
	userAccess := []model.UserRole{ model.Regular, model.Admin }
	adminAccess := []model.UserRole{ model.Admin }

	roleMap["/proto.OrderService/GetOrders"] = regularAccess
	roleMap["/proto.OrderService/CountOrders"] = regularAccess
	roleMap["/proto.OrderService/AddOrders"] = userAccess
	roleMap["/proto.OrderService/EditOrders"] = adminAccess
	roleMap["/proto.OrderService/DeleteOrders"] = adminAccess


	return roleMap
}

func (h *Handler) ProvideAuthKey() *rsa.PublicKey {
	return h.auth.PublicKey()
}