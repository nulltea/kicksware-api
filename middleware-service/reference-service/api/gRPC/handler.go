package gRPC

import (
	"reference-service/api/gRPC/proto"
	"reference-service/core/meta"
	"reference-service/core/model"
	"reference-service/core/service"
	"reference-service/env"
)

//go:generate protoc --go_out=plugins=grpc:. proto/reference.proto

type Handler struct {
	service     service.SneakerReferenceService
	auth        service.AuthService
	contentType string
}

func NewHandler(service service.SneakerReferenceService, auth service.AuthService, config env.CommonConfig) *Handler {
	return &Handler{
		service,
		auth,
		config.ContentType,
	}
}

func (h *Handler) GetReferences(filter *proto.ReferenceFilter, srv proto.ReferenceService_GetReferencesServer) (err error) {
	var references []*model.SneakerReference
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}

	if len(filter.ReferenceID) == 0 && filter.RequestQuery == nil  {
		references, err = h.service.FetchAll(params)
	} else if filter.RequestQuery != nil {
		query, _ := meta.NewRequestQuery(filter.RequestQuery)
		references, err = h.service.FetchQuery(query, params)
	} else if len(filter.ReferenceID) == 1 {
		ref, e := h.service.FetchOne(filter.ReferenceID[0], params); if e != nil {
			err = e
		}
		references = []*model.SneakerReference {ref}
	} else {
		references, err = h.service.Fetch(filter.ReferenceID, params)
	}

	srv.Send(&proto.ReferenceResponse{
		References: proto.NativeToReferences(references),
	 	Count: int64(len(references)),
	})
	return
}

func (h *Handler) CountReferences(filter *proto.ReferenceFilter, srv proto.ReferenceService_CountReferencesServer) (err error) {
	var count int = 0

	if filter == nil {
		count, err = h.service.CountAll()
	} else if filter.RequestQuery != nil {
		query, _ := meta.NewRequestQuery(filter.RequestQuery)
		count, err = h.service.Count(query, filter.RequestParams.ToNative())
	}

	go srv.Send(&proto.ReferenceResponse{
		References: nil,
		Count: int64(count),
	})
	return
}

func (h *Handler) AddReferences(input *proto.ReferenceInput, srv proto.ReferenceService_AddReferencesServer) (err error) {
	var succeeded int64
	var references []*model.SneakerReference

	for _, user := range input.References {
		native := user.ToNative()
		if err := h.service.StoreOne(native); err != nil {
			succeeded = succeeded + 1
			references = append(references, native)
		}
	}

	go srv.Send(&proto.ReferenceResponse{
		References: proto.NativeToReferences(references),
		Count: succeeded,
	})
	return
}

func (h *Handler) EditReferences(input *proto.ReferenceInput, srv proto.ReferenceService_EditReferencesServer) (err error) {
	var succeeded int64

	for _, user := range input.References {
		if err := h.service.Modify(user.ToNative()); err != nil {
			succeeded = succeeded + 1
		}
	}

	go srv.Send(&proto.ReferenceResponse{Count: succeeded})
	return
}