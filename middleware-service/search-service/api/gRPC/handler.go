package gRPC

import (
	"context"

	"search-service/api/gRPC/proto"
	"search-service/core/meta"
	"search-service/core/model"
	"search-service/core/service"
	"search-service/env"
)

//go:generate protoc --proto_path=../../../service-protos --go_out=plugins=grpc:proto/. reference.proto
//go:generate protoc --proto_path=../../../service-protos --go_out=plugins=grpc:proto/. product.proto
//go:generate protoc --proto_path=../../../service-protos  --go_out=plugins=grpc:proto/. common.proto
//go:generate protoc --proto_path=../../../service-protos --go_out=plugins=grpc:proto/. search.proto


type Handler struct {
	search      service.ReferenceSearchService
	sync        service.ReferenceSyncService
	auth        service.AuthService
	contentType string
}

func NewHandler(search service.ReferenceSearchService, sync service.ReferenceSyncService, auth service.AuthService, config env.CommonConfig) *Handler {
	return &Handler{
		search,
		sync,
		auth,
		config.ContentType,
	}
}

func (h *Handler) Search(ctx context.Context, tag *proto.SearchTag) (resp *proto.ReferenceResponse, err error) {
	refs, err :=  h.search.Search(tag.Tag, tag.RequestParams.ToNative()); if err != nil {
		return
	}
	resp = &proto.ReferenceResponse{
		References: proto.NativeToReferences(refs),
		Count: int64(len(refs)),
	}
	return
}

func (h *Handler) SearchBy(ctx context.Context, filter *proto.SearchFilter) (resp *proto.ReferenceResponse, err error) {
	var refs []*model.SneakerReference
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}

	if len(filter.SKU) > 0 {
		refs, err =  h.search.SearchSKU(filter.SKU, params); if err != nil {
			return
		}
	} else if len(filter.Model) > 0 {
		refs, err =  h.search.SearchModel(filter.Model, params); if err != nil {
			return
		}
	} else if len(filter.Brand) > 0 {
		refs, err =  h.search.SearchBrand(filter.Brand, params); if err != nil {
			return
		}
	} else {
		refs, err =  h.search.SearchBy(filter.Field, filter.Value, params); if err != nil {
			return
		}
	}

	resp = &proto.ReferenceResponse{
		References: proto.NativeToReferences(refs),
		Count: int64(len(refs)),
	}
	return
}

func (h *Handler) Sync(ctx context.Context, filter *proto.ReferenceFilter) (resp *proto.ReferenceResponse, err error) {
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}

	if len(filter.ReferenceID) == 0 && filter.RequestQuery == nil  {
		err = h.sync.SyncAll(params)
	} else if filter.RequestQuery != nil {
		query, _ := meta.NewRequestQuery(filter.RequestQuery)
		err = h.sync.SyncQuery(query, params)
	} else if len(filter.ReferenceID) == 1 {
		err = h.sync.SyncOne(filter.ReferenceID[0])
	} else {
		err = h.sync.Sync(filter.ReferenceID, params)
	}
	resp = &proto.ReferenceResponse{}
	return
}
