package gRPC

import (
	"crypto/rsa"

	"go.kicksware.com/api/shared/core/meta"
	"google.golang.org/grpc"

	"go.kicksware.com/api/services/users/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterUserServiceServer(server, handler)
		proto.RegisterAuthServiceServer(server, handler)
		proto.RegisterMailServiceServer(server, handler)
		proto.RegisterInteractServiceServer(server, handler)
	}
}

func (h *Handler) ProvideAccessRoles() meta.AccessConfig {
	return meta.AccessConfig {
		"/proto.UserService/GetUsers": meta.RegularAccess,
		"/proto.UserService/CountUsers": meta.RegularAccess,
		"/proto.UserService/AddUsers": meta.RegularAccess,
		"/proto.UserService/EditUsers": meta.UserAccess,
		"/proto.UserService/DeleteUsers": meta.AdminAccess,

		"/proto.AuthService/SignUp": meta.GuestAccess,
		"/proto.AuthService/Login": meta.GuestAccess,
		"/proto.AuthService/Guest": meta.ZeroAccess,
		"/proto.AuthService/GenerateToken": meta.UserAccess,
		"/proto.AuthService/Refresh": meta.RegularAccess,
		"/proto.AuthService/Logout": meta.UserAccess,

		"/proto.MailService/SendEmailConfirmation": meta.RegularAccess,
		"/proto.MailService/SendResetPassword": meta.RegularAccess,
		"/proto.MailService/SendNotification": meta.AdminAccess,
		"/proto.MailService/Subscribe": meta.RegularAccess,
		"/proto.MailService/Unsubscribe": meta.RegularAccess,

		"/proto.InteractService/Like": meta.UserAccess,
		"/proto.InteractService/Unlike": meta.UserAccess,
	}
}

func (h *Handler) ProvideAuthKey() *rsa.PublicKey {
	return h.auth.PublicKey()
}
