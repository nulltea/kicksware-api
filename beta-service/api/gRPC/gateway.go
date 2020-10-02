package grpc

import (
	"crypto/rsa"

	"github.com/timoth-y/kicksware-api/user-service/core/model"
	gRPC "google.golang.org/grpc"

	"go.kicksware.com/api/beta-service/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *gRPC.Server) {
	return func(server *gRPC.Server) {
		proto.RegisterBetaServiceServer(server, handler)
	}
}

func (h *Handler) ProvideAccessRoles() map[string][]model.UserRole {
	roleMap := make(map[string][]model.UserRole)

	var zeroAccess []model.UserRole
	guestAccess := []model.UserRole{ model.Guest }
	regularAccess := []model.UserRole{ model.Guest, model.Regular, model.Admin }
	userAccess := []model.UserRole{ model.Regular, model.Admin }
	adminAccess := []model.UserRole{ model.Admin }

	roleMap["/proto.BetaService/GetBetas"] = zeroAccess
	roleMap["/proto.BetaService/CountBetas"] = guestAccess
	roleMap["/proto.BetaService/AddBetas"] = regularAccess
	roleMap["/proto.BetaService/EditBetas"] = userAccess
	roleMap["/proto.BetaService/DeleteBetas"] = adminAccess


	return roleMap
}

func (h *Handler) ProvideAuthKey() *rsa.PublicKey {
	return h.auth.PublicKey()
}