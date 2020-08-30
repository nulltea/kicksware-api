package meta

import (
	"strings"

	"github.com/timoth-y/kicksware-platform/middleware-service/beta-service/core/model"
)

type RequestParams struct {
	limit int
	offset int
	sortBy string
	sortDirection string
}

func (p *RequestParams) Limit() int {
	return p.limit
}
func (p *RequestParams) SetLimit(limit int) {
	p.limit = limit
}

func (p *RequestParams) Offset() int {
	return p.offset
}
func (p *RequestParams) SetOffset(offset int) {
	p.offset = offset
}

func (p *RequestParams) SortBy() string {
	return strings.ToLower(p.sortBy)
}
func (p *RequestParams) SetSortBy(sortBy string) {
	p.sortBy = sortBy
}

func (p *RequestParams) SortDirection() string {
	return p.sortDirection
}
func (p *RequestParams) SortDirectionNum() int {
	if p.sortDirection == "desc" {
		return -1
	}
	return 1
}
func (p *RequestParams) SetSortDirection(direction string) {
	p.sortDirection = direction
}

func (p *RequestParams) ApplyParams(users []*model.Beta) []*model.Beta {
	if p.sortBy != "" {
		// business.NewSorter(users, p.sortBy).Sort(p.sortDirection == "desc")
	}
	if p.offset != 0 {
		users = users[p.offset:]
	}
	if p.limit != 0 && p.limit < len(users) {
		users = users[:p.limit]
	}
	return users
}