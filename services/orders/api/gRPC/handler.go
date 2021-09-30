package gRPC

import (
	"context"

	"github.com/pkg/errors"
	"go.kicksware.com/api/shared/core"
	"go.kicksware.com/api/shared/util"

	"go.kicksware.com/api/shared/core/meta"

	"go.kicksware.com/api/services/orders/api/gRPC/proto"
	"go.kicksware.com/api/services/orders/core/model"
	"go.kicksware.com/api/services/orders/core/service"
	"go.kicksware.com/api/services/orders/usecase/business"
)

//go:generate protoc --proto_path=../../../service-protos --go_out=plugins=grpc,paths=source_relative:proto/. orders.proto

type Handler struct {
	service service.OrderService
	auth    core.AuthService
}

func NewHandler(service service.OrderService, auth core.AuthService) *Handler {
	return &Handler{
		service,
		auth,
	}
}

func (h *Handler) GetOrders(ctx context.Context, filter *proto.OrderFilter) (resp *proto.OrderResponse, err error) {
	var orders []*model.Order
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}
	util.SetAuthParamsFromMetaData(ctx, &params)

	if len(filter.OrderID) == 0 && filter.RequestQuery == nil {
		orders, err = h.service.FetchAll(params)
	} else if filter.RequestQuery != nil {
		query := filter.RequestQuery.AsMap()
		orders, err = h.service.FetchQuery(query, params)
	} else if len(filter.OrderID) == 1 {
		ref, e := h.service.FetchOne(filter.OrderID[0], params); if e != nil {
			err = e
		}
		orders = []*model.Order{ref}
	} else {
		orders, err = h.service.Fetch(filter.OrderID, params)
	}

	if errors.Cause(err) == business.ErrOrderNotFound {
		return &proto.OrderResponse{
			Count: 0,
		}, nil
	}

	resp = &proto.OrderResponse{
		Orders: proto.NativeToOrders(orders),
	 	Count:  int64(len(orders)),
	}
	return
}

func (h *Handler) CountOrders(ctx context.Context, filter *proto.OrderFilter) (resp *proto.OrderResponse, err error) {
	var count int = 0
	var params *meta.RequestParams; if filter != nil && filter.RequestParams != nil {
		params = filter.RequestParams.ToNative()
	}
	util.SetAuthParamsFromMetaData(ctx, &params)

	if len(filter.OrderID) == 0 && filter.RequestQuery == nil {
		count, err = h.service.CountAll()
	} else if filter.RequestQuery != nil {
		query := filter.RequestQuery.AsMap()
		count, err = h.service.Count(query, params)
	}

	resp = &proto.OrderResponse{
		Orders: nil,
		Count: int64(count),
	}
	return
}

func (h *Handler) AddOrder(ctx context.Context, input *proto.OrderInput) (resp *proto.OrderResponse, err error) {
	userID, ok := getInputUserID(ctx, input); if !ok {
		return nil, business.ErrOrderNotValid
	}

	order := &model.Order{
		UserID: userID,
		ReferenceID: input.ReferenceID,
		ProductID: input.ProductID,
	}

	if err = h.service.StoreOne(order); err != nil {
		return nil, err
	}

	resp = &proto.OrderResponse{
		Orders: []*proto.Order{
			(&proto.Order{}).FromNative(order),
		}, Count: 1,
	}
	return
}

func (h *Handler) EditOrder(ctx context.Context, input *proto.OrderInput) (resp *proto.OrderResponse, err error) {
	userID, ok := getInputUserID(ctx, input); if !ok {
		return nil, business.ErrOrderNotValid
	}
	if input.Order == nil || userID != input.Order.UserID {
		return nil, business.ErrOrderNotValid
	}

	order := input.Order.ToNative()
	if err := h.service.Modify(order); err != nil {
		return nil, err
	}

	resp = &proto.OrderResponse{
		Orders: []*proto.Order{
			(&proto.Order{}).FromNative(order),
		}, Count: 1,
	}
	return
}

func (h *Handler) DeleteOrder(ctx context.Context, filter *proto.OrderFilter) (resp *proto.OrderResponse, err error) {
	orderIDs := filter.OrderID; orderID := orderIDs[0]; if len(orderIDs) == 0 || len(orderID) == 0 {
		return nil, business.ErrOrderNotFound
	}

	order, err := h.service.FetchOne(orderID, nil); if err != nil {
		return nil, business.ErrOrderNotFound
	}

	if userID, ok := util.RetrieveUserID(ctx); !ok {
		return nil, business.ErrOrderNotValid
	} else if userID != order.UserID {
		return nil, business.ErrOrderNotValid
	}

	if err = h.service.Remove(orderID); err != nil {
		return nil, err
	}

	resp = &proto.OrderResponse{
		Orders: nil, Count: 1,
	}
	return
}

func getInputUserID(ctx context.Context, input *proto.OrderInput) (string, bool) {
	userID := input.UserID
	if id, ok := util.RetrieveUserID(ctx); ok {
		userID = id
	}; if len(userID) == 0 {
		return "", false
	}
	return userID, true
}
