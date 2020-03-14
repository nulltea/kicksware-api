package tool

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
)

func ToMap(v interface{}) (m map[string]interface{}, err error) {
	data, err := json.Marshal(v)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &m)
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
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &m)
	return
}
