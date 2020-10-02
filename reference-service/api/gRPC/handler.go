package gRPC

import (
	"context"

	"github.com/pkg/errors"
	"github.com/timoth-y/kicksware-api/service-common/config"
	"github.com/timoth-y/kicksware-api/service-common/core"
	"github.com/timoth-y/kicksware-api/service-common/util"

	"github.com/timoth-y/kicksware-api/service-common/core/meta"

	"github.com/timoth-y/kicksware-api/reference-service/api/gRPC/proto"
	"github.com/timoth-y/kicksware-api/reference-service/core/model"
	"github.com/timoth-y/kicksware-api/reference-service/core/service"
	"github.com/timoth-y/kicksware-api/reference-service/usecase/business"
)

//go:generate protoc --proto_path=../../../service-protos --go_out=plugins=grpc,paths=source_relative:proto/. reference.proto

type Handler struct {
	service     service.SneakerReferenceService
	auth        core.AuthService
	contentType string
}

func NewHandler(service service.SneakerReferenceService, auth core.AuthService, config config.CommonConfig) *Handler {
	return &Handler{
		service,
		auth,
		config.ContentType,
	}
}

func (h *Handler) GetReferences(ctx context.Context, filter *proto.ReferenceFilter) (resp *proto.ReferenceResponse, err error) {
	var references []*model.SneakerReference
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}
	util.SetAuthParamsFromMetaData(ctx, &params)

	if len(filter.ReferenceID) == 0 && filter.RequestQuery == nil {
		references, err = h.service.FetchAll(params)
	} else if filter.RequestQuery != nil {
		query := filter.RequestQuery.AsMap()
		references, err = h.service.FetchQuery(query, params)
	} else if len(filter.ReferenceID) == 1 {
		ref, e := h.service.FetchOne(filter.ReferenceID[0], params); if e != nil {
			err = e
		}
		references = []*model.SneakerReference {ref}
	} else {
		references, err = h.service.Fetch(filter.ReferenceID, params)
	}

	if errors.Cause(err) == business.ErrReferenceNotFound {
		return &proto.ReferenceResponse{
			Count: 0,
		}, nil
	}

	resp = &proto.ReferenceResponse{
		References: proto.NativeToReferences(references),
	 	Count: int64(len(references)),
	}
	return
}

func (h *Handler) CountReferences(ctx context.Context, filter *proto.ReferenceFilter) (resp *proto.ReferenceResponse, err error) {
	var count int = 0
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}
	util.SetAuthParamsFromMetaData(ctx, &params)

	if len(filter.ReferenceID) == 0 && filter.RequestQuery == nil {
		count, err = h.service.CountAll()
	} else if filter.RequestQuery != nil {
		query := filter.RequestQuery.AsMap()
		count, err = h.service.Count(query, params)
	}

	resp = &proto.ReferenceResponse{
		References: nil,
		Count: int64(count),
	}
	return
}

func (h *Handler) AddReferences(ctx context.Context, input *proto.ReferenceInput) (resp *proto.ReferenceResponse, err error) {
	var succeeded int64
	var references []*model.SneakerReference

	for _, user := range input.References {
		native := user.ToNative()
		if err := h.service.StoreOne(native); err != nil {
			succeeded = succeeded + 1
			references = append(references, native)
		}
	}

	resp = &proto.ReferenceResponse{
		References: proto.NativeToReferences(references),
		Count: succeeded,
	}
	return
}

func (h *Handler) EditReferences(ctx context.Context, input *proto.ReferenceInput) (resp *proto.ReferenceResponse, err error) {
	var succeeded int64

	for _, user := range input.References {
		if err := h.service.Modify(user.ToNative()); err != nil {
			succeeded = succeeded + 1
		}
	}

	resp = &proto.ReferenceResponse{Count: succeeded}
	return
}

func (h *Handler) DeleteReferences(ctx context.Context, filter *proto.ReferenceFilter) (resp *proto.ReferenceResponse, err error) {
	panic("implement me")
}
