package business

import (
	"sort"
	"strings"

	"go.kicksware.com/api/services/references/core/model"
)

type ReferenceSorter struct {
	items []*model.SneakerReference
	property string
}

func NewSorter(items []*model.SneakerReference, property string) (s *ReferenceSorter) {
	s = &ReferenceSorter{}
	s.items = items
	s.property = property
	return s
}
func (s *ReferenceSorter) Len() int      { return len(s.items) }
func (s *ReferenceSorter) Swap(i, j int) { s.items[i], s.items[j] = s.items[j], s.items[i] }
func (s *ReferenceSorter) Less(i, j int) bool {
	switch strings.ToLower(s.property) {
	case "price":
		return s.items[i].Price < s.items[j].Price
	case "released":
		return s.items[i].ReleaseDate.Sub(s.items[j].ReleaseDate).Hours() < 0
	default:
		return s.items[i].UniqueId < s.items[j].UniqueId
	}
}
func (s *ReferenceSorter) Asc() []*model.SneakerReference {
	sort.Sort(s)
	return s.items
}

func (s *ReferenceSorter) Desc() []*model.SneakerReference {
	sort.Sort(sort.Reverse(s))
	return s.items
}

func (s *ReferenceSorter) Sort(desc bool) []*model.SneakerReference {
	if desc  {
		return s.Desc()
	}
	return s.Asc();
}

