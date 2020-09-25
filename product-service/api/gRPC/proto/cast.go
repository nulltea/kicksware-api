package proto

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/timoth-y/kicksware-api/product-service/core/meta"
	"github.com/timoth-y/kicksware-api/product-service/core/model"
)

func (m *SneakerProduct) ToNative() *model.SneakerProduct {
	return &model.SneakerProduct{
		UniqueId:       m.UniqueId,
		BrandName:      m.BrandName,
		ModelName:      m.ModelName,
		ModelSKU:       m.ModelSKU,
		ReferenceId:    m.ReferenceId,
		Price:          m.Price,
		Type:           m.Type,
		Size:           m.Size.ToNative(),
		Color:          m.Color,
		Condition:      m.Condition,
		Description:    m.Description,
		Owner:          m.Owner,
		// Images:         m.Images,
		ConditionIndex: m.ConditionIndex,
		AddedAt:        m.AddedAt.AsTime(),
	}
}

func (m *SneakerProduct) FromNative(n *model.SneakerProduct) *SneakerProduct {
	m.UniqueId = n.UniqueId
	m.BrandName = n.BrandName
	m.ModelName = n.ModelName
	m.ModelSKU = n.ModelSKU
	m.ReferenceId = n.ReferenceId
	m.Price = n.Price
	m.Type = n.Type
	m.Size = SneakerSize{}.FromNative(n.Size)
	m.Color = n.Color
	m.Condition = n.Condition
	m.Description = n.Description
	m.Owner = n.Owner
	// m.Images         = n.Images
	m.ConditionIndex = n.ConditionIndex
	m.AddedAt = timestamppb.New(n.AddedAt)
	return m
}

func (m *SneakerSize) ToNative() model.SneakerSize {
	return model.SneakerSize{
		Europe:        m.Europe,
		UnitedStates:  m.UnitedStates,
		UnitedKingdom: m.UnitedKingdom,
		Centimeters:   m.Centimeters,
	}
}

func (m SneakerSize) FromNative(n model.SneakerSize) *SneakerSize {
	m.Europe = n.Europe
	m.UnitedStates = n.UnitedStates
	m.UnitedKingdom = n.UnitedKingdom
	m.Centimeters = n.Centimeters
	return &m
}

func NativeToProducts(native []*model.SneakerProduct) []*SneakerProduct {
	users := make([]*SneakerProduct, 0)
	for _, user := range native {
		users = append(users, (&SneakerProduct{}).FromNative(user))
	}
	return users
}

func ProductsToNative(in []*SneakerProduct) []*model.SneakerProduct {
	users := make([]*model.SneakerProduct, 0)
	for _, user := range in {
		users = append(users, user.ToNative())
	}
	return users
}

func (m *RequestParams) ToNative() *meta.RequestParams {
	n := &meta.RequestParams{}
	n.SetLimit(int(m.Limit))
	n.SetOffset(int(m.Offset))
	if m.SortBy != nil {
		n.SetSortBy(m.SortBy.Value)
	}
	if m.SortDirection != nil {
		n.SetSortDirection(m.SortDirection.Value)
	}
	return n
}

func (m RequestParams) FromNative(n *meta.RequestParams) *RequestParams {
	m.Limit = int32(n.Limit())
	m.Offset = int32(n.Offset())
	m.SortBy = &wrappers.StringValue{Value: n.SortBy()}
	m.SortDirection = &wrappers.StringValue{Value: n.SortDirection()}
	return &m
}