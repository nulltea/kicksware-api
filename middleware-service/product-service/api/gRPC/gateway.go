package gRPC

import (
	"google.golang.org/grpc"

	"product-service/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterProductServiceServer(server, handler)
	}
}