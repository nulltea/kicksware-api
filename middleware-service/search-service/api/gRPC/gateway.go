package gRPC

import (
	"google.golang.org/grpc"

	"search-service/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterSearchReferencesServiceServer(server, handler)
	}
}