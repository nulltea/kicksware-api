package gRPC

import (
	"crypto/rsa"

	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/model"
	"google.golang.org/grpc"

	"github.com/timoth-y/kicksware-platform/middleware-service/reference-service/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterReferenceServiceServer(server, handler)
	}
}

func (h *Handler) ProvideAccessRoles() map[string][]model.UserRole {
	roleMap := make(map[string][]model.UserRole)

	regularAccess := []model.UserRole{ model.Guest, model.Regular, model.Admin }
	userAccess := []model.UserRole{ model.Regular, model.Admin }
	adminAccess := []model.UserRole{ model.Admin }

	roleMap["/proto.ReferenceService/GetReferences"] = regularAccess
	roleMap["/proto.ReferenceService/CountReferences"] = regularAccess
	roleMap["/proto.ReferenceService/AddReferences"] = userAccess
	roleMap["/proto.ReferenceService/EditReferences"] = adminAccess
	roleMap["/proto.ReferenceService/DeleteReferences"] = adminAccess


	return roleMap
}

func (h *Handler) ProvideAuthKey() *rsa.PublicKey {
	return h.auth.PublicKey()
}