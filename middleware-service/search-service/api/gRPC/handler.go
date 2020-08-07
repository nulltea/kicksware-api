package gRPC

import (
	"context"

	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/api/gRPC/proto"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/meta"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/model"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/service"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/env"
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
	var params *meta.RequestParams; if tag != nil && tag.RequestParams != nil {
		params = tag.RequestParams.ToNative()
	}

	refs, err :=  h.search.Search(tag.Tag, params); if err != nil {
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

	refs, err =  h.search.SearchBy(filter.Field, filter.Value, params); if err != nil {
		return
	}

	resp = &proto.ReferenceResponse{
		References: proto.NativeToReferences(refs),
		Count: int64(len(refs)),
	}
	return
}

func (h *Handler) SearchSKU(ctx context.Context, filter *proto.SearchFilter) (resp *proto.ReferenceResponse, err error) {
	var refs []*model.SneakerReference
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}

	refs, err =  h.search.SearchSKU(filter.Value, params); if err != nil {
		return
	}

	resp = &proto.ReferenceResponse{
		References: proto.NativeToReferences(refs),
		Count: int64(len(refs)),
	}
	return
}

func (h *Handler) SearchBrand(ctx context.Context, filter *proto.SearchFilter) (resp *proto.ReferenceResponse, err error) {
	var refs []*model.SneakerReference
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}

	refs, err =  h.search.SearchBrand(filter.Value, params); if err != nil {
		return
	}

	resp = &proto.ReferenceResponse{
		References: proto.NativeToReferences(refs),
		Count: int64(len(refs)),
	}
	return
}

func (h *Handler) SearchModel(ctx context.Context, filter *proto.SearchFilter) (resp *proto.ReferenceResponse, err error) {
	var refs []*model.SneakerReference
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}

	refs, err =  h.search.SearchModel(filter.Value, params); if err != nil {
		return
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
