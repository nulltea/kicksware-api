package common

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/fatih/structs"

	"reference-service/core/model"
	"reference-service/middleware/business"
)

type RequestParams struct {
	TakeCount int
	SkipOffset int
	SortBy string
	SortDirection string
}

func (p *RequestParams) ApplyParams(references []*model.SneakerReference) []*model.SneakerReference {
	if p.SortBy != "" {
		business.NewSorter(references, p.SortBy).Sort(p.SortDirection == "desc")
	}
	if p.SkipOffset != 0 {
		references = references[p.SkipOffset:]
	}
	if p.TakeCount != 0 && p.TakeCount < len(references) {
		references = references[:p.TakeCount]
	}
	return references
}

func (p *RequestParams) AssignParams(r *http.Request) {
	query := r.URL.Query();
	properties := structs.Names(p)
	for _, prop := range properties {
		value := query.Get(prop)
		if value == "" {
			continue
		}
		field := reflect.ValueOf(p).Elem().FieldByName(prop);
		switch field.Kind().String() {
		case "string":
			field.SetString(value);
		case "int", "float":
			if num, err := strconv.ParseInt(value, 10, 32); err == nil {
				field.SetInt(num);
			}
		case "bool":
			if sign, err := strconv.ParseBool(value); err == nil {
				field.SetBool(sign);
			}
		default:
			field.SetString(value);
		}
	}
	return
}