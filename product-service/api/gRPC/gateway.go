package gRPC

import (
	"crypto/rsa"

	"go.kicksware.com/api/user-service/core/model"
	"google.golang.org/grpc"

	"go.kicksware.com/api/product-service/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterProductServiceServer(server, handler)
	}
}

func (h *Handler) ProvideAccessRoles() map[string][]model.UserRole {
	roleMap := make(map[string][]model.UserRole)

	regularAccess := []model.UserRole{ model.Guest, model.Regular, model.Admin }
	userAccess := []model.UserRole{ model.Regular, model.Admin }

	roleMap["/proto.ProductService/GetProducts"] = regularAccess
	roleMap["/proto.ProductService/CountProducts"] = regularAccess
	roleMap["/proto.ProductService/AddProducts"] = userAccess
	roleMap["/proto.ProductService/EditProducts"] = userAccess
	roleMap["/proto.ProductService/DeleteProducts"] = userAccess


	return roleMap
}

func (h *Handler) ProvideAuthKey() *rsa.PublicKey {
	return h.auth.PublicKey()
}