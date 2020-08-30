package meta

import (
	sqb "github.com/Masterminds/squirrel"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestQuery map[string]interface{}

func NewRequestQuery(i interface{}) (RequestQuery, error) {
	return toMap(i)
}

func (q RequestQuery) ToBson() (m bson.M, err error) {
	return toBson(q)
}

func (q RequestQuery) ToSql() (sqb.And, error) {
	dict, err := toMap(q)
	if err != nil {
		return nil, err
	}
	keys := funk.Keys(dict).([]string)
	cond := funk.Map(keys, func(k interface{}) sqb.Sqlizer {
		key := k.(string)
		return sqb.Eq{key:dict[key]}
	}).([]sqb.Sqlizer)
	res := append(sqb.And{}, cond...)
	return res, nil
}

func toBson(v interface{}) (m bson.M, err error) {
	dict, err := toMap(v)
	if err != nil {
		return
	}
	m = bson.M {}
	for key := range dict {
		switch dict[key].(type) {
		case map[string]interface{}:
			sub, err := toBson(dict[key])
			if err != nil {
				m[key] = dict[key]
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
		case []interface{}, primitive.A:
			values := bson.A{}
			for _, item := range dict[key].(primitive.A) {
				switch item.(type) {
				case map[string]interface{}:
					sub, err := toBson(item)
					if err != nil {
						continue
					}
					values = append(values, sub)
				default:
					values = append(values, item)
				}
			}
			m[key] = values
		default:
			m[key] = dict[key]
		}
	}
	return
}

func toMap(v interface{}) (m map[string]interface{}, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &m)
	return
}