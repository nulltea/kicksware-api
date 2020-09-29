package business

import (
	"sort"
	"strings"

	"github.com/timoth-y/kicksware-api/order-service/core/model"
)

type OrderSorter struct {
	items []*model.Order
	property string
}

func NewSorter(items []*model.Order, property string) (s *OrderSorter) {
	s = &OrderSorter{}
	s.items = items
	s.property = property
	return s
}
func (s *OrderSorter) Len() int           { return len(s.items) }
func (s *OrderSorter) Swap(i, j int)      { s.items[i], s.items[j] = s.items[j], s.items[i] }
func (s *OrderSorter) Less(i, j int) bool {
	switch strings.ToLower(s.property) {
	case "price":
		return s.items[i].Price < s.items[j].Price
	case "released":
		return s.items[i].OrderedAt.Sub(s.items[j].OrderedAt).Hours() < 0
	default:
		return s.items[i].UniqueID < s.items[j].UniqueID
	}
}
func (s *OrderSorter) Asc() []*model.Order {
	sort.Sort(s)
	return s.items
}

func (s *OrderSorter) Desc() []*model.Order {
	sort.Sort(sort.Reverse(s))
	return s.items
}

func (s *OrderSorter) Sort(desc bool) []*model.Order {
	if desc  {
		return s.Desc()
	}
	return s.Asc();
}

