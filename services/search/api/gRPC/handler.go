package gRPC

import (
	"context"

	protoRef "go.kicksware.com/api/services/references/api/gRPC/proto"
	"go.kicksware.com/api/services/references/core/model"
	"go.kicksware.com/api/shared/core"
	"go.kicksware.com/api/shared/core/meta"
	"go.kicksware.com/api/shared/util"

	"go.kicksware.com/api/services/search/api/gRPC/proto"
	"go.kicksware.com/api/services/search/core/service"
)

//go:generate protoc --proto_path=../../../service-protos --go_out=plugins=grpc,paths=source_relative:proto/. search.proto

type Handler struct {
	search service.ReferenceSearchService
	sync   service.ReferenceSyncService
	auth   core.AuthService
}

func NewHandler(search service.ReferenceSearchService, sync service.ReferenceSyncService, auth core.AuthService) *Handler {
	return &Handler{
		search,
		sync,
		auth,
	}
}

func (h *Handler) Search(ctx context.Context, tag *proto.SearchTag) (resp *protoRef.ReferenceResponse, err error) {
	var params *meta.RequestParams; if tag != nil && tag.RequestParams != nil {
		params = tag.RequestParams.ToNative()
	}
	util.SetAuthParamsFromMetaData(ctx, &params)

	refs, err :=  h.search.Search(tag.Tag, params); if err != nil {
		return
	}
	resp = &protoRef.ReferenceResponse{
		References: protoRef.NativeToReferences(refs),
		Count: int64(len(refs)),
	}
	return
}

func (h *Handler) SearchBy(ctx context.Context, filter *proto.SearchFilter) (resp *protoRef.ReferenceResponse, err error) {
	var refs []*model.SneakerReference
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}
	util.SetAuthParamsFromMetaData(ctx, &params)

	refs, err =  h.search.SearchBy(filter.Field, filter.Value, params); if err != nil {
		return
	}

	resp = &protoRef.ReferenceResponse{
		References: protoRef.NativeToReferences(refs),
		Count: int64(len(refs)),
	}
	return
}

func (h *Handler) SearchSKU(ctx context.Context, filter *proto.SearchFilter) (resp *protoRef.ReferenceResponse, err error) {
	var refs []*model.SneakerReference
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}
	util.SetAuthParamsFromMetaData(ctx, &params)

	refs, err =  h.search.SearchSKU(filter.Value, params); if err != nil {
		return
	}

	resp = &protoRef.ReferenceResponse{
		References: protoRef.NativeToReferences(refs),
		Count: int64(len(refs)),
	}
	return
}

func (h *Handler) SearchBrand(ctx context.Context, filter *proto.SearchFilter) (resp *protoRef.ReferenceResponse, err error) {
	var refs []*model.SneakerReference
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}
	util.SetAuthParamsFromMetaData(ctx, &params)

	refs, err =  h.search.SearchBrand(filter.Value, params); if err != nil {
		return
	}

	resp = &protoRef.ReferenceResponse{
		References: protoRef.NativeToReferences(refs),
		Count: int64(len(refs)),
	}
	return
}

func (h *Handler) SearchModel(ctx context.Context, filter *proto.SearchFilter) (resp *protoRef.ReferenceResponse, err error) {
	var refs []*model.SneakerReference
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}
	util.SetAuthParamsFromMetaData(ctx, &params)

	refs, err =  h.search.SearchModel(filter.Value, params); if err != nil {
		return
	}

	resp = &protoRef.ReferenceResponse{
		References: protoRef.NativeToReferences(refs),
		Count: int64(len(refs)),
	}
	return
}

func (h *Handler) Sync(ctx context.Context, filter *protoRef.ReferenceFilter) (resp *protoRef.ReferenceResponse, err error) {
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}
	util.SetAuthParamsFromMetaData(ctx, &params)

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
	resp = &protoRef.ReferenceResponse{}
	return
}
