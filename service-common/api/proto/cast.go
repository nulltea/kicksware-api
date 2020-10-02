package proto

import (
	"github.com/golang/protobuf/ptypes/wrappers"

	"go.kicksware.com/api/service-common/core/meta"
)

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