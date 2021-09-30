package gRPC

import (
	"context"

	"github.com/golang/glog"
	"github.com/pkg/errors"
	"go.kicksware.com/api/shared/core"
	"go.kicksware.com/api/shared/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.kicksware.com/api/shared/core/meta"

	"go.kicksware.com/api/services/products/api/gRPC/proto"
	"go.kicksware.com/api/services/products/core/model"
	"go.kicksware.com/api/services/products/core/service"
	"go.kicksware.com/api/services/products/usecase/business"
)

//go:generate protoc --proto_path=../../../service-protos --go_out=plugins=grpc,paths=source_relative:proto/. product.proto

type Handler struct {
	service service.SneakerProductService
	auth    core.AuthService
}

func NewHandler(service service.SneakerProductService, auth core.AuthService) *Handler {
	return &Handler{
		service,
		auth,
	}
}

func (h *Handler) GetProducts(ctx context.Context, filter *proto.ProductFilter) (resp *proto.ProductResponse, err error) {
	var products []*model.SneakerProduct
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}
	util.SetAuthParamsFromMetaData(ctx, &params)

	if len(filter.ProductID) == 0 && filter.RequestQuery == nil  {
		products, err = h.service.FetchAll(params)
	} else if filter.RequestQuery != nil {
		query, _ := meta.NewRequestQuery(filter.RequestQuery)
		products, err = h.service.FetchQuery(query, params)
	} else if len(filter.ProductID) == 1 {
		user, e := h.service.FetchOne(filter.ProductID[0]); if e != nil {
			err = e
		}
		products = []*model.SneakerProduct{user}
	} else {
		products, err = h.service.Fetch(filter.ProductID, params)
	}

	if errors.Cause(err) == business.ErrProductNotFound {
		return &proto.ProductResponse{
			Count: 0,
		}, nil
	}

	resp = &proto.ProductResponse{
		Products: proto.NativeToProducts(products),
		Count:    int64(len(products)),
	}
	return
}

func (h *Handler) CountProducts(ctx context.Context, filter *proto.ProductFilter) (resp *proto.ProductResponse, err error) {
	var count int = 0
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}
	util.SetAuthParamsFromMetaData(ctx, &params)

	if filter == nil {
		count, err = h.service.CountAll()
	} else if filter.RequestQuery != nil {
		query, _ := meta.NewRequestQuery(filter.RequestQuery)
		count, err = h.service.Count(query, params)
	}

	resp = &proto.ProductResponse{
		Products: nil,
		Count: int64(count),
	}
	return
}

func (h *Handler) AddProducts(ctx context.Context, input *proto.ProductInput) (resp *proto.ProductResponse, err error) {
	var succeeded int64
	var products []*model.SneakerProduct

	for _, user := range input.Products {
		native := user.ToNative()
		if err := h.service.Store(native, nil); err != nil {
			succeeded = succeeded + 1
			products = append(products, native)
		}
	}

	resp = &proto.ProductResponse{
		Products: proto.NativeToProducts(products),
		Count:    succeeded,
	}
	return
}

func (h *Handler) EditProducts(ctx context.Context, input *proto.ProductInput) (resp *proto.ProductResponse, err error) {
	var succeeded int64

	for _, user := range input.Products {
		if err := h.service.Modify(user.ToNative()); err != nil {
			succeeded = succeeded + 1
		}
	}

	resp = &proto.ProductResponse{Count: succeeded}
	return
}

func (h *Handler) DeleteProducts(ctx context.Context, filter *proto.ProductFilter) (resp *proto.ProductResponse, err error) {
	var count int = 0

	if len(filter.ProductID) == 0 && filter.RequestQuery == nil {
		glog.Errorln("Could not delete all products")
	} else if filter.RequestQuery != nil {
		glog.Errorln("Could not delete products by condition")
	} else if len(filter.ProductID) == 1 {
		e := h.service.Remove(filter.ProductID[0]); if e != nil {
			err = e
		}
		count = 1
	} else {
		glog.Errorln("Could not delete all reference")
	}

	resp = &proto.ProductResponse{
		Count: int64(count),
	}
	return
}

func (h *Handler) UploadImages(ctx context.Context, input *proto.UploadImageRequest) (*proto.ProductResponse, error) {
	productID := input.ProductID; if len(input.ProductID) == 0 {
		return nil, status.Error(codes.InvalidArgument, business.ErrProductNotFound.Error())
	}

	if len(input.Images) == 0 {
		return nil, status.Error(codes.InvalidArgument, "images data must be sent with request")
	}

	product, err := h.service.FetchOne(productID); if err != nil {
		return nil, status.Error(codes.InvalidArgument, business.ErrProductNotFound.Error())
	}

	if product.Images == nil || len(product.Images) == 0 {
		product.Images = input.Images
	} else {
		for key, data := range input.Images {
			product.Images[key] = data
		}
	}

	err = h.service.Modify(product); if err != nil {
		return nil, status.Errorf(codes.Internal,
			"Internal error occurred while modifying product record: %q", err,
		)
	}
	return &proto.ProductResponse{
		Count: 1,
	}, nil
}

func (h *Handler) RequestAnalysis(ctx context.Context, input *proto.ProductInput) (*proto.AnalysisResponse, error) {
	panic("implement me")
}

func (h *Handler) RequestPrediction(ctx context.Context, request *proto.PredictionRequest) (*proto.ProductResponse, error) {
	panic("implement me")
}
