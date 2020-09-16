package business

import (
	"sort"
	"strings"

	"github.com/timoth-y/kicksware-api/user-service/core/model"
)

type UserSorter struct {
	items []*model.User
	property string
}

func NewSorter(items []*model.User, property string) (s *UserSorter) {
	s = &UserSorter{}
	s.items = items
	s.property = property
	return s
}
func (s *UserSorter) Len() int      { return len(s.items) }
func (s *UserSorter) Swap(i, j int) { s.items[i], s.items[j] = s.items[j], s.items[i] }
func (s *UserSorter) Less(i, j int) bool {
	switch strings.ToLower(s.property) {
	case "registered":
		return s.items[i].RegisterDate.Sub(s.items[j].RegisterDate).Hours() < 0
	default:
		return s.items[i].UniqueID < s.items[j].UniqueID
	}
}
func (s *UserSorter) Asc() []*model.User {
	sort.Sort(s)
	return s.items
}

func (s *UserSorter) Desc() []*model.User {
	sort.Sort(sort.Reverse(s))
	return s.items
}

func (s *UserSorter) Sort(desc bool) []*model.User {
	if desc  {
		return s.Desc()
	}
	return s.Asc();
}

