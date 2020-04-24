package util

import (
	"net/url"

	sqb "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToMap(v interface{}) map[string]interface{} {
	return structs.Map(v)
}

func toMap(v interface{}) (m map[string]interface{}, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &m)
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

func ToBsonMap(v interface{}) (m bson.M, err error) {
	data, err := toMap(v)
	if err != nil {
		return
	}
	m = bson.M {}
	for key := range data {
		switch data[key].(type) {
		case map[string]interface{}:
			sub, err := ToBsonMap(data[key])
			if err != nil {
				m[key] = data[key]
			} else {
				if key == "$regex" {
					m[key] = primitive.Regex{
						Pattern: sub["pattern"].(string),
						Options: sub["options"].(string),
					}
				} else {
					m[key] = sub
				}
			}
		default:
			m[key] = data[key]
		}
	}
	return
}

func ToSqlWhere(v interface{}) sqb.And {
	filter := ToMap(v)
	keys := funk.Keys(filter).([]string)
	cond := funk.Map(keys, func(k interface{}) sqb.Sqlizer {
		key := k.(string)
		return sqb.Eq{key:filter[key]}
	}).([]sqb.Sqlizer)
	res := append(sqb.And{}, cond...)
	return res
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