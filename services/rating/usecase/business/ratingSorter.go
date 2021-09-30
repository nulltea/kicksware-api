package business

import (
	"sort"
	"strings"

	"go.kicksware.com/api/services/rating/core/model"
)

type RatingSorter struct {
	items []*model.Rating
	property string
}

func NewSorter(items []*model.Rating, property string) (s *RatingSorter) {
	s = &RatingSorter{}
	s.items = items
	s.property = property
	return s
}
func (s *RatingSorter) Len() int      { return len(s.items) }
func (s *RatingSorter) Swap(i, j int) { s.items[i], s.items[j] = s.items[j], s.items[i] }
func (s *RatingSorter) Less(i, j int) bool {
	switch strings.ToLower(s.property) {
	case "rating":
		return s.items[i].Rating < s.items[j].Rating
	case "views":
		return s.items[i].Views < s.items[j].Views
	case "searches":
		return s.items[i].Searches < s.items[j].Searches
	case "orders":
		return s.items[i].Orders < s.items[j].Orders
	default:
		return s.items[i].UniqueID < s.items[j].UniqueID
	}
}
func (s *RatingSorter) Asc() []*model.Rating {
	sort.Sort(s)
	return s.items
}

func (s *RatingSorter) Desc() []*model.Rating {
	sort.Sort(sort.Reverse(s))
	return s.items
}

func (s *RatingSorter) Sort(desc bool) []*model.Rating {
	if desc  {
		return s.Desc()
	}
	return s.Asc()
}

