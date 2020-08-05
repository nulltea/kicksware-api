package grpc

import (
	"context"

	"beta-service/api/gRPC/proto"
	"beta-service/core/meta"
	"beta-service/core/model"
	"beta-service/core/service"
	"beta-service/env"
)

//go:generate protoc --go_out=plugins=grpc:. proto/beta.proto

type Handler struct {
	service     service.BetaService
	auth        service.AuthService
	contentType string
}

func NewHandler(service service.BetaService, auth service.AuthService, config env.CommonConfig) *Handler {
	return &Handler{
		service,
		auth,
		config.ContentType,
	}
}

func (h* Handler) GetBetas(ctx context.Context, filter *proto.BetaFilter) (r *proto.BetaResponse, err error) {
	var users []*model.Beta

	if filter == nil || (len(filter.BetaID) == 0 && filter.RequestQuery == nil) {
		users, err = h.service.FetchAll(nil)
	} else if filter.RequestQuery != nil {
		query, _ := meta.NewRequestQuery(filter.RequestQuery)
		users, err = h.service.FetchQuery(query, filter.RequestParams.ToNative())
	} else if len(filter.BetaID) == 1 {
		user, e := h.service.FetchOne(filter.BetaID[0], nil); if e != nil {
			err = e
		}
		users = []*model.Beta {user}
	} else {
		users, err = h.service.Fetch(filter.BetaID, filter.RequestParams.ToNative())
	}

	r = &proto.BetaResponse{
		Betas: proto.NativeToBetas(users),
		Count: int64(len(users)),
	}
	return
}

func (h* Handler) CountBetas(ctx context.Context, filter *proto.BetaFilter) (r *proto.BetaResponse, err error) {
	var count int = 0

	if filter == nil {
		count, err = h.service.CountAll()
	} else if filter.RequestQuery != nil {
		query, _ := meta.NewRequestQuery(filter.RequestQuery)
		count, err = h.service.Count(query, filter.RequestParams.ToNative())
	}

	r = &proto.BetaResponse{
		Betas: nil,
		Count: int64(count),
	}
	return
}

func (h* Handler) AddBetas(ctx context.Context, input *proto.BetaInput) (*proto.BetaResponse, error) {
	var succeeded int64
	var users []*model.Beta

	for _, user := range input.Betas {
		native := user.ToNative()
		if err := h.service.StoreOne(native); err != nil {
			succeeded = succeeded + 1
			users = append(users, native)
		}
	}

	return &proto.BetaResponse{
		Betas: proto.NativeToBetas(users),
		Count: succeeded,
	}, nil
}

func (h* Handler) EditBetas(ctx context.Context, input *proto.BetaInput) (*proto.BetaResponse, error) {
	var succeeded int64

	for _, user := range input.Betas {
		if err := h.service.Modify(user.ToNative()); err != nil {
			succeeded = succeeded + 1
		}
	}

	return &proto.BetaResponse{Count: succeeded}, nil
}
