package gRPC

import (
	"google.golang.org/grpc"

	"user-service/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterUserServiceServer(server, handler)
		proto.RegisterAuthServiceServer(server, handler)
		proto.RegisterMailServiceServer(server, handler)
		proto.RegisterInteractServiceServer(server, handler)
	}
}