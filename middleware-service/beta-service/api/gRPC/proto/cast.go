package proto

import (
	"github.com/timoth-y/kicksware-platform/middleware-service/beta-service/core/meta"
	"github.com/timoth-y/kicksware-platform/middleware-service/beta-service/core/model"
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
	n.SetSortBy(m.SortBy)
	n.SetSortDirection(m.SortDirection)
	return n
}