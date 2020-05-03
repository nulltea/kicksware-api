package util

import (
	"net/url"

	"github.com/fatih/structs"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
)

func ToMap(v interface{}) map[string]interface{} {
	return structs.Map(v)
}

func ToBsonMap(v interface{}) (d bson.M, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &d)
	return
}

func ToBsonDoc(v interface{}) (d bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &d)
	return
}

func GetAllInsertValues(v interface{}) []interface{} {
	return structs.Values(v)
}

func GetAllInsertColumns(v interface{}) []string {
	return structs.Names(v)
}

func GetInsertValues(v interface{}, fields []string) []interface{} {
	filter  := ToMap(v)
	values := funk.Map(fields, func(k interface{}) interface{}{
		key := k.(string)
		return filter[key]
	}).([]interface{})
	return values
}

func ToQueryMap(v url.Values) (qm map[string]interface{}) {
	qm = make(map[string]interface{})
	keys := funk.Keys(v).([]string)
	for _, key := range keys {
		if len(v[key]) > 1 {
			qm[key] = v[key]
			continue
		}
		qm[key] = v[key][0]
	}
	return
}