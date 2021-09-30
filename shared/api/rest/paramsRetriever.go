package rest

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/fatih/structs"

	"go.kicksware.com/api/shared/core/meta"
)

func NewRequestParams(r *http.Request) *meta.RequestParams {
	p := &meta.RequestParams{}
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
		p.SetUserID(r.URL.User.Username())
		if token, ok := r.URL.User.Password(); ok {
			p.SetToken(token)
		}
	}
	return p
}


func setPrivateField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).
		Elem().
		Set(reflect.ValueOf(value))
}
