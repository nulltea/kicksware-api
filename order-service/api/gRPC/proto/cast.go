package proto

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/timoth-y/kicksware-api/order-service/core/model"
)

func (m *Order) ToNative() *model.Order {
	return &model.Order{
		UniqueID:    m.UniqueID,
		UserID:      m.UserID,
		ReferenceID: m.ReferenceID,
		ProductID:   m.ProductID,
		Price:       m.Price,
		Status:      model.OrderStatus(m.Status),
		SourceURL:   m.SourceURL,
		AddedAt:     m.AddedAt.AsTime(),
	}
}

func (m *Order) FromNative(n *model.Order) *Order {
	m.UniqueID = n.UniqueID
	m.UserID = n.UserID
	m.ReferenceID = n.ReferenceID
	m.ProductID = n.ProductID
	m.Price = n.Price
	m.Status = string(n.Status)
	m.SourceURL = n.SourceURL
	m.AddedAt = timestamppb.New(n.AddedAt)
	return m
}

func NativeToOrders(native []*model.Order) []*Order {
	users := make([]*Order, 0)
	for _, user := range native {
		users = append(users, (&Order{}).FromNative(user))
	}
	return users
}

func OrdersToNative(in []*Order) []*model.Order {
	users := make([]*model.Order, 0)
	for _, user := range in {
		users = append(users, user.ToNative())
	}
	return users
}