package proto

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/timoth-y/kicksware-api/reference-service/core/meta"
	"github.com/timoth-y/kicksware-api/reference-service/core/model"
)

func (m *SneakerReference) ToNative() *model.SneakerReference {
	return &model.SneakerReference{
		UniqueId:       m.UniqueId,
		ManufactureSku: m.ManufactureSku,
		BrandName:      m.BrandName,
		Brand:          m.Brand.ToNative(),
		ModelName:      m.ModelName,
		Model:          m.Model.ToNative(),
		BaseModelName:  m.BaseModelName,
		BaseModel:      m.BaseModel.ToNative(),
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
		Likes:          int(m.Likes),
		Liked:          m.Liked,
	}
}

func (m *SneakerReference) FromNative(n *model.SneakerReference) *SneakerReference {
	m.UniqueId = n.UniqueId
	m.UniqueId = n.UniqueId
	m.ManufactureSku = n.ManufactureSku
	m.BrandName = n.BrandName
	m.Brand = SneakerBrand{}.FromNative(n.Brand)
	m.ModelName = n.ModelName
	m.Model = SneakerModel{}.FromNative(n.Model)
	m.BaseModelName = n.BaseModelName
	m.BaseModel = SneakerModel{}.FromNative(n.Model)
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

func (m *SneakerBrand) ToNative() model.SneakerBrand {
	return model.SneakerBrand{
		UniqueId:    m.UniqueId,
		Name:        m.Name,
		Logo:        m.Logo,
		Hero:        m.Hero,
		Description: m.Description,
	}
}

func (m SneakerBrand) FromNative(n model.SneakerBrand) *SneakerBrand {
	m.UniqueId = n.UniqueId
	m.Name = n.Name
	m.Logo = n.Logo
	m.Hero = n.Hero
	m.Description = n.Description
	return &m
}

func (m *SneakerModel) ToNative() model.SneakerModel {
	return model.SneakerModel{
		UniqueId:    m.UniqueId,
		Name:        m.Name,
		Brand:       m.Brand,
		BaseModel:   m.BaseModel,
		Hero:        m.Hero,
		Description: m.Description,
	}
}

func (m SneakerModel) FromNative(n model.SneakerModel) *SneakerModel {
	m.UniqueId = n.UniqueId
	m.Name = n.Name
	m.Brand = n.Brand
	m.BaseModel = n.BaseModel
	m.Hero = n.Hero
	m.Description = n.Description
	return &m
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

func (m *RequestParams) ToNative() *meta.RequestParams {
	n := &meta.RequestParams{}
	n.SetLimit(int(m.Limit))
	n.SetOffset(int(m.Offset))
	if m.SortBy != nil {
		n.SetSortBy(m.SortBy.Value)
	}
	if n.SortDirection != nil {
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