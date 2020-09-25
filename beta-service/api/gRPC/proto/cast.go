package proto

import (
	"github.com/golang/protobuf/ptypes/wrappers"

	"github.com/timoth-y/kicksware-api/beta-service/core/meta"
	"github.com/timoth-y/kicksware-api/beta-service/core/model"
)

func (m *Beta) ToNative() *model.Beta {
	return &model.Beta{
		UniqueID: m.UniqueID,
		Feature:  m.Feature,
		Done:     m.Done,
	}
}

func (m *Beta) FromNative(n *model.Beta) *Beta {
	m.UniqueID = n.UniqueID
	m.Feature = n.Feature
	m.Done = n.Done
	return m
}

func NativeToBetas(native []*model.Beta) []*Beta {
	users := make([]*Beta, 0)
	for _, user := range native {
		users = append(users, (&Beta{}).FromNative(user))
	}
	return users
}

func BetasToNative(in []*Beta) []*model.Beta {
	users := make([]*model.Beta, 0)
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