package gRPC

import (
	"context"

	"search-service/api/gRPC/proto"
	"search-service/core/meta"
	"search-service/core/model"
	"search-service/core/service"
	"search-service/env"
)

//go:generate protoc --go_out=plugins=grpc:. proto/search.proto

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

func (h *Handler) Search(tag *proto.SearchTag, srv proto.SearchReferencesService_SearchServer) error {
	refs, err :=  h.search.Search(tag.Tag, tag.RequestParams.ToNative()); if err != nil {
		return err
	}
	srv.Send(&proto.ReferenceResponse{
		References: proto.NativeToReferences(refs),
		Count: int64(len(refs)),
	})
	return nil
}

func (h *Handler) SearchBy(filter *proto.SearchFilter, srv proto.SearchReferencesService_SearchByServer) (err error){
	var refs []*model.SneakerReference
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}

	if len(filter.SKU) > 0 {
		refs, err =  h.search.SearchSKU(filter.SKU, params); if err != nil {
			return err
		}
	} else if len(filter.Model) > 0 {
		refs, err =  h.search.SearchModel(filter.Model, params); if err != nil {
			return err
		}
	} else if len(filter.Brand) > 0 {
		refs, err =  h.search.SearchBrand(filter.Brand, params); if err != nil {
			return err
		}
	} else {
		refs, err =  h.search.SearchBy(filter.Field, filter.Value, params); if err != nil {
			return err
		}
	}

	srv.Send(&proto.ReferenceResponse{
		References: proto.NativeToReferences(refs),
		Count: int64(len(refs)),
	})
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
