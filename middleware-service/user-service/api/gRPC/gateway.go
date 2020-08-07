package gRPC

import (
	"crypto/rsa"

	"google.golang.org/grpc"

	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/api/gRPC/proto"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/model"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterUserServiceServer(server, handler)
		proto.RegisterAuthServiceServer(server, handler)
		proto.RegisterMailServiceServer(server, handler)
		proto.RegisterInteractServiceServer(server, handler)
	}
}

func (h *Handler) ProvideAccessRoles() map[string][]model.UserRole {
	roleMap := make(map[string][]model.UserRole)

	var zeroAccess []model.UserRole
	guestAccess := []model.UserRole{ model.Guest }
	regularAccess := []model.UserRole{ model.Guest, model.Regular, model.Admin }
	userAccess := []model.UserRole{ model.Regular, model.Admin }
	adminAccess := []model.UserRole{ model.Admin }

	roleMap["/proto.UserService/GetUsers"] = regularAccess
	roleMap["/proto.UserService/CountUsers"] = regularAccess
	roleMap["/proto.UserService/AddUsers"] = regularAccess
	roleMap["/proto.UserService/EditUsers"] = userAccess
	roleMap["/proto.UserService/DeleteUsers"] = adminAccess

	roleMap["/proto.AuthService/SignUp"] = guestAccess
	roleMap["/proto.AuthService/Login"] = zeroAccess
	roleMap["/proto.AuthService/Guest"] = zeroAccess
	roleMap["/proto.AuthService/GenerateToken"] = userAccess
	roleMap["/proto.AuthService/Refresh"] = regularAccess
	roleMap["/proto.AuthService/Logout"] = userAccess

	roleMap["/proto.MailService/SendEmailConfirmation"] = regularAccess
	roleMap["/proto.MailService/SendResetPassword"] = regularAccess
	roleMap["/proto.MailService/SendNotification"] = adminAccess

	roleMap["/proto.InteractService/Like"] = userAccess
	roleMap["/proto.InteractService/Unlike"] = userAccess

	return roleMap
}

func (h *Handler) ProvideAuthKey() *rsa.PublicKey {
	return h.auth.PublicKey()
}