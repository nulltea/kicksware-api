package gRPC

//go:generate protoc --go_out=plugins=grpc:. proto/user.proto

import (
	"context"

	"user-service/api/gRPC/proto"
	"user-service/core/meta"
	"user-service/core/model"
	"user-service/core/service"
	"user-service/env"
)

type Handler struct {
	service     service.UserService
	auth        service.AuthService
	mail        service.MailService
	interact    service.InteractService
	contentType string
}

func NewHandler(service service.UserService, auth service.AuthService, mail service.MailService,
	interact service.InteractService, config env.CommonConfig) *Handler {
	return &Handler{
		service,
		auth,
		mail,
		interact,
		config.ContentType,
	}
}

func (h* Handler) GetUsers(ctx context.Context, filter *proto.UserFilter) (r *proto.UserResponse, err error) {
	var users []*model.User
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}

	if len(filter.UserID) == 0 && filter.RequestQuery == nil  {
		users, err = h.service.FetchAll(params)
	} else if filter.RequestQuery != nil {
		query, _ := meta.NewRequestQuery(filter.RequestQuery)
		users, err = h.service.FetchQuery(query, params)
	} else if len(filter.UserID) == 1 {
		user, e := h.service.FetchOne(filter.UserID[0]); if e != nil {
			err = e
		}
		users = []*model.User {user}
	} else {
		users, err = h.service.Fetch(filter.UserID, params)
	}

	r = &proto.UserResponse{
		Users: proto.NativeToUsers(users),
		Count: int64(len(users)),
	}
	return
}

func (h* Handler) CountUsers(ctx context.Context, filter *proto.UserFilter) (r *proto.UserResponse, err error) {
	var count int = 0

	if filter == nil {
		count, err = h.service.CountAll()
	} else if filter.RequestQuery != nil {
		query, _ := meta.NewRequestQuery(filter.RequestQuery)
		count, err = h.service.Count(query, filter.RequestParams.ToNative())
	}

	r = &proto.UserResponse{
		Users: nil,
		Count: int64(count),
	}
	return
}

func (h* Handler) AddUsers(ctx context.Context, input *proto.UserInput) (*proto.UserResponse, error) {
	var succeeded int64
	var users []*model.User

	for _, user := range input.Users {
		native := user.ToNative()
		if err := h.service.Register(native); err != nil {
			succeeded = succeeded + 1
			users = append(users, native)
		}
	}

	return &proto.UserResponse{
		Users: proto.NativeToUsers(users),
		Count: succeeded,
	}, nil
}

func (h* Handler) EditUsers(ctx context.Context, input *proto.UserInput) (*proto.UserResponse, error) {
	var succeeded int64

	for _, user := range input.Users {
		if err := h.service.Modify(user.ToNative()); err != nil {
			succeeded = succeeded + 1
		}
	}

	return &proto.UserResponse{Count: succeeded}, nil
}

