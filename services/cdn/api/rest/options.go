package rest

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/fatih/structs"

	"go.kicksware.com/api/services/cdn/core/meta"
)

func ParseOptions(r *http.Request) meta.ImageOptions {
	o := &meta.ImageOptions{}
	query := r.URL.Query()
	properties := structs.Names(o)
	for _, prop := range properties {
		value := query.Get(strings.ToLower(prop))
		if value == "" {
			continue
		}
		field := reflect.ValueOf(o).Elem().FieldByName(prop)
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
	return *o
}

func setPrivateField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).
		Elem().
		Set(reflect.ValueOf(value))
}
