package gRPC

import (
	"crypto/rsa"

	"go.kicksware.com/api/service-common/core/meta"
	"google.golang.org/grpc"

	"go.kicksware.com/api/search-service/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterSearchReferencesServiceServer(server, handler)
	}
}

func (h *Handler) ProvideAccessRoles() meta.AccessConfig {
	return meta.AccessConfig{
		"/proto.SearchReferencesService/Search": meta.RegularAccess,
		"/proto.SearchReferencesService/SearchBy": meta.RegularAccess,
		"/proto.SearchReferencesService/SearchSKU": meta.RegularAccess,
		"/proto.SearchReferencesService/SearchBrand": meta.RegularAccess,
		"/proto.SearchReferencesService/SearchModel": meta.RegularAccess,
		"/proto.SearchReferencesService/Sync": meta.AdminAccess,

		"/proto.SearchProductService/Search": meta.RegularAccess,
		"/proto.SearchProductService/SearchBy": meta.RegularAccess,
		"/proto.SearchProductService/Sync": meta.AdminAccess,
	}
}

func (h *Handler) ProvideAuthKey() *rsa.PublicKey {
	return h.auth.PublicKey()
}

