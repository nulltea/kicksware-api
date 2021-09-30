package gRPC

import (
	"bytes"
	"context"
	"errors"

	"go.kicksware.com/api/shared/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.kicksware.com/api/services/cdn/api/gRPC/proto"
	"go.kicksware.com/api/services/cdn/core/service"
)

//go:generate protoc --proto_path=../../../service-protos --go_out=plugins=grpc,paths=source_relative:proto/. cdn.proto

var (
	ErrQueryInvalid = errors.New("content query invalid")
)

type Handler struct {
	service service.ContentService
	auth    core.AuthService
}

func NewHandler(service service.ContentService, auth core.AuthService) *Handler {
	return &Handler{
		service,
		auth,
	}
}

func (h *Handler) Original(ctx context.Context, request *proto.ContentRequest) (*proto.Content, error) {
	query := request.ToNative(); if len(query.Collection) == 0 || len(query.Filename) == 0 {
		return nil, status.Error(codes.InvalidArgument, ErrQueryInvalid.Error())
	}
	content, err := h.service.Original(query); if err != nil {
		return nil, err
	}
	return proto.Content{}.FromNative(content), nil
}

func (h *Handler) Crop(ctx context.Context, request *proto.ContentRequest) (*proto.Content, error) {
	query := request.ToNative(); if len(query.Collection) == 0 || len(query.Filename) == 0 ||
		query.ImageOptions.Width == 0 && query.ImageOptions.Height == 0 {
		return nil, status.Error(codes.InvalidArgument, ErrQueryInvalid.Error())
	}
	content, err := h.service.Crop(query, query.ImageOptions); if err != nil {
		return nil, err
	}
	return proto.Content{}.FromNative(content), nil
}

func (h *Handler) Resize(ctx context.Context, request *proto.ContentRequest) (*proto.Content, error) {
	query := request.ToNative(); if len(query.Collection) == 0 || len(query.Filename) == 0 ||
		query.ImageOptions.Width == 0 && query.ImageOptions.Height == 0 {
		return nil, status.Error(codes.InvalidArgument, ErrQueryInvalid.Error())
	}
	content, err := h.service.Resize(query, query.ImageOptions); if err != nil {
		return nil, err
	}
	return proto.Content{}.FromNative(content), nil
}

func (h *Handler) Thumbnail(ctx context.Context, request *proto.ContentRequest) (*proto.Content, error) {
	query := request.ToNative(); if len(query.Collection) == 0 || len(query.Filename) == 0 {
		return nil, status.Error(codes.InvalidArgument, ErrQueryInvalid.Error())
	}
	content, err := h.service.Thumbnail(query); if err != nil {
		return nil, err
	}
	return proto.Content{}.FromNative(content), nil
}


func (h *Handler) Upload(server proto.ContentService_UploadServer) error {
	input, err := server.Recv(); if err != nil {
		return err
	}
	request := input.Request.ToNative(); if len(input.Data) == 0 || len(request.Collection) == 0 || len(request.Filename) == 0 {
		return status.Error(codes.InvalidArgument, ErrQueryInvalid.Error())
	}
	if err := h.service.Upload(bytes.NewBuffer(input.Data), request); err != nil {
		return err
	}
	return err
}

func (h *Handler) StreamContent(request *proto.ContentRequest, server proto.ContentService_StreamContentServer) error {
	query := request.ToNative(); if len(query.Collection) == 0 || len(query.Filename) == 0 {
		return status.Error(codes.InvalidArgument, ErrQueryInvalid.Error())
	}
	content, err := h.service.Original(query); if err != nil {
		return err
	}
	server.Send(proto.Content{}.FromNative(content))
	return nil
}
