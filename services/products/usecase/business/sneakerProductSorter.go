package business

import (
	"sort"
	"strings"

	"go.kicksware.com/api/services/products/core/model"
)

type ProductSorter struct {
	items []*model.SneakerProduct
	property string
}

func NewSorter(items []*model.SneakerProduct, property string) (s *ProductSorter) {
	s = &ProductSorter{}
	s.items = items
	s.property = property
	return s
}
func (s *ProductSorter) Len() int      { return len(s.items) }
func (s *ProductSorter) Swap(i, j int) { s.items[i], s.items[j] = s.items[j], s.items[i] }
func (s *ProductSorter) Less(i, j int) bool {
	switch strings.ToLower(s.property) {
	case "price":
		return s.items[i].Price < s.items[j].Price
	case "added":
		return s.items[i].AddedAt.Sub(s.items[j].AddedAt).Hours() < 0
	default:
		return s.items[i].UniqueId < s.items[j].UniqueId
	}
}
func (s *ProductSorter) Asc() []*model.SneakerProduct {
	sort.Sort(s)
	return s.items
}

func (s *ProductSorter) Desc() []*model.SneakerProduct {
	sort.Sort(sort.Reverse(s))
	return s.items
}

func (s *ProductSorter) Sort(desc bool) []*model.SneakerProduct {
	if desc  {
		return s.Desc()
	}
	return s.Asc();
}

