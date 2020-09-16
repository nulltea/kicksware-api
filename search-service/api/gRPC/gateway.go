package gRPC

import (
	"crypto/rsa"

	"github.com/timoth-y/kicksware-api/user-service/core/model"
	"google.golang.org/grpc"

	"github.com/timoth-y/kicksware-api/search-service/api/gRPC/proto"
)

func ProvideRemoteSetup(handler *Handler) func(server *grpc.Server) {
	return func(server *grpc.Server) {
		proto.RegisterSearchReferencesServiceServer(server, handler)
	}
}

func (h *Handler) ProvideAccessRoles() map[string][]model.UserRole {
	roleMap := make(map[string][]model.UserRole)

	regularAccess := []model.UserRole{ model.Guest, model.Regular, model.Admin }
	adminAccess := []model.UserRole{ model.Admin }

	roleMap["/proto.SearchReferencesService/Search"] = regularAccess
	roleMap["/proto.SearchReferencesService/SearchBy"] = regularAccess
	roleMap["/proto.SearchReferencesService/SearchSKU"] = regularAccess
	roleMap["/proto.SearchReferencesService/SearchBrand"] = regularAccess
	roleMap["/proto.SearchReferencesService/SearchModel"] = regularAccess
	roleMap["/proto.SearchReferencesService/Sync"] = adminAccess

	roleMap["/proto.SearchProductService/Search"] = regularAccess
	roleMap["/proto.SearchProductService/SearchBy"] = regularAccess
	roleMap["/proto.SearchProductService/Sync"] = adminAccess

	return roleMap
}

func (h *Handler) ProvideAuthKey() *rsa.PublicKey {
	return h.auth.PublicKey()
}