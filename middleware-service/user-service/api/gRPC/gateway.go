package grpc

import (
	gRPC "google.golang.org/grpc"

	"user-service/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *gRPC.Server) {
	return func(server *gRPC.Server) {
		proto.RegisterUserServiceServer(server, handler)
	}
}