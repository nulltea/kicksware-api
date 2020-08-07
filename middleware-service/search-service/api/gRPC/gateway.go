package gRPC

import (
	"google.golang.org/grpc"

	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterSearchReferencesServiceServer(server, handler)
	}
}