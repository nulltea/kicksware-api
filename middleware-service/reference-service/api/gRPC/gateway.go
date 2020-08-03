package gRPC

import (
	"google.golang.org/grpc"

	"reference-service/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterReferenceServiceServer(server, handler)
	}
}