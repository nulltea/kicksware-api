package gRPC

import (
	"product-service/api/gRPC/proto"
	"product-service/core/meta"
	"product-service/core/model"
	"product-service/core/service"
	"product-service/env"
)

//go:generate protoc --go_out=plugins=grpc:. proto/product.proto

type Handler struct {
	service     service.SneakerProductService
	auth        service.AuthService
	contentType string
}

func NewHandler(service service.SneakerProductService, auth service.AuthService, config env.CommonConfig) *Handler {
	return &Handler{
		service,
		auth,
		config.ContentType,
	}
}

func (h *Handler) GetProducts(filter *proto.ProductFilter, srv proto.ProductService_GetProductsServer) (err error) {
	var products []*model.SneakerProduct
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}

	if len(filter.ProductID) == 0 && filter.RequestQuery == nil  {
		products, err = h.service.FetchAll(params)
	} else if filter.RequestQuery != nil {
		query, _ := meta.NewRequestQuery(filter.RequestQuery)
		products, err = h.service.FetchQuery(query, params)
	} else if len(filter.ProductID) == 1 {
		user, e := h.service.FetchOne(filter.ProductID[0]); if e != nil {
			err = e
		}
		products = []*model.SneakerProduct {user}
	} else {
		products, err = h.service.Fetch(filter.ProductID, params)
	}

	srv.Send(&proto.ProductResponse{
		Products: proto.NativeToProducts(products),
		Count: int64(len(products)),
	})
	return
}

func (h *Handler) CountProducts(filter *proto.ProductFilter, srv proto.ProductService_CountProductsServer) (err error) {
	var count int = 0

	if filter == nil {
		count, err = h.service.CountAll()
	} else if filter.RequestQuery != nil {
		query, _ := meta.NewRequestQuery(filter.RequestQuery)
		count, err = h.service.Count(query, filter.RequestParams.ToNative())
	}

	go srv.Send(&proto.ProductResponse{
		Products: nil,
		Count: int64(count),
	})
	return
}

func (h *Handler) AddProducts(input *proto.ProductInput, srv proto.ProductService_AddProductsServer) (err error) {
	var succeeded int64
	var products []*model.SneakerProduct

	for _, user := range input.Products {
		native := user.ToNative()
		if err := h.service.Store(native, nil); err != nil {
			succeeded = succeeded + 1
			products = append(products, native)
		}
	}

	go srv.Send(&proto.ProductResponse{
		Products: proto.NativeToProducts(products),
		Count: succeeded,
	})
	return
}

func (h *Handler) EditProducts(input *proto.ProductInput, srv proto.ProductService_EditProductsServer) (err error) {
	var succeeded int64

	for _, user := range input.Products {
		if err := h.service.Modify(user.ToNative()); err != nil {
			succeeded = succeeded + 1
		}
	}

	go srv.Send(&proto.ProductResponse{Count: succeeded})
	return
}