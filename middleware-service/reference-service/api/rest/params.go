package rest

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/fatih/structs"

	"reference-service/core/meta"
	"reference-service/core/model"
	"reference-service/usecase/business"
)

type params struct {
	limit int
	offset int
	sortBy string
	sortDirection string
	userID string
}

func NewRequestParams(r *http.Request) meta.RequestParams {
	p := &params{}
	query := r.URL.Query()
	properties := structs.Names(p)
	for _, prop := range properties {
		value := query.Get(strings.ToLower(prop))
		if value == "" {
			continue
		}
		field := reflect.ValueOf(p).Elem().FieldByName(prop)
		switch field.Kind().String() {
		case "string":
			setPrivateField(field, value)
		case "int", "float":
			if num, err := strconv.ParseInt(value, 10, 32); err == nil {
				setPrivateField(field, int(num))
			}
		case "bool":
			if sign, err := strconv.ParseBool(value); err == nil {
				setPrivateField(field, sign)
			}
		default:
			setPrivateField(field, value)
		}
	}
	if r.URL.User != nil {
		p.userID = r.URL.User.Username()
	}

	return p
}

func (p *params) Limit() int {
	return p.limit
}
func (p *params) SetLimit(limit int) {
	p.limit = limit
}

func (p *params) Offset() int {
	return p.offset
}
func (p *params) SetOffset(offset int) {
	p.offset = offset
}

func (p *params) SortBy() string {
	return strings.ToLower(p.sortBy)
}
func (p *params) SetSortBy(sortBy string) {
	p.sortBy = sortBy
}

func (p *params) SortDirection() string {
	return p.sortDirection
}
func (p *params) SortDirectionNum() int {
	if p.sortDirection == "desc" {
		return -1
	}
	return 1
}
func (p *params) SetSortDirection(direction string) {
	p.sortDirection = direction
}

func (p *params) UserID() string {
	return p.userID
}

func (p *params) SetUserID(userID string) {
	p.userID = userID
}

func (p *params) ApplyParams(references []*model.SneakerReference) []*model.SneakerReference {
	if p.sortBy != "" {
		business.NewSorter(references, p.sortBy).Sort(p.sortDirection == "desc")
	}
	if p.offset != 0 {
		references = references[p.offset:]
	}
	if p.limit != 0 && p.limit < len(references) {
		references = references[:p.limit]
	}
	return references
}

func setPrivateField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).
		Elem().
		Set(reflect.ValueOf(value))
}