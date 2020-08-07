package proto

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/meta"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/model"
)

func (m *SneakerReference) ToNative() *model.SneakerReference {
	return &model.SneakerReference{
		UniqueId:       m.UniqueId,
		ManufactureSku: m.ManufactureSku,
		BrandName:      m.BrandName,
		ModelName:      m.ModelName,
		BaseModelName:  m.BaseModelName,
		Description:    m.Description,
		Color:          m.Color,
		Gender:         m.Gender,
		Nickname:       m.Nickname,
		Materials:      m.Materials,
		Categories:     m.Categories,
		ReleaseDate:    m.ReleaseDate.AsTime(),
		Price:          m.Price,
		ImageLink:      m.ImageLink,
		ImageLinks:     m.ImageLinks,
		StadiumUrl:     m.StadiumUrl,
	}
}

func (m *SneakerReference) FromNative(n *model.SneakerReference) *SneakerReference {
	m.UniqueId = n.UniqueId
	m.UniqueId = n.UniqueId
	m.ManufactureSku = n.ManufactureSku
	m.BrandName = n.BrandName
	m.ModelName = n.ModelName
	m.BaseModelName = n.BaseModelName
	m.Description = n.Description
	m.Color = n.Color
	m.Gender = n.Gender
	m.Nickname = n.Nickname
	m.Materials = n.Materials
	m.Categories = n.Categories
	m.ReleaseDate = timestamppb.New(n.ReleaseDate)
	m.Price = n.Price
	m.ImageLink = n.ImageLink
	m.ImageLinks = n.ImageLinks
	m.StadiumUrl = n.StadiumUrl
	return m
}

func NativeToReferences(native []*model.SneakerReference) []*SneakerReference {
	users := make([]*SneakerReference, 0)
	for _, user := range native {
		users = append(users, (&SneakerReference{}).FromNative(user))
	}
	return users
}

func ReferencesToNative(in []*SneakerReference) []*model.SneakerReference {
	users := make([]*model.SneakerReference, 0)
	for _, user := range in {
		users = append(users, user.ToNative())
	}
	return users
}

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
	n.SetSortBy(m.SortBy)
	n.SetSortDirection(m.SortDirection)
	return n
}

func (m RequestParams) FromNative(n *meta.RequestParams) *RequestParams {
	m.Limit = int32(n.Limit())
	m.Offset = int32(n.Offset())
	m.SortBy = n.SortBy()
	m.SortDirection = n.SortDirection()
	return &m
}
