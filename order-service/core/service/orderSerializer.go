package service

import model "go.kicksware.com/api/order-service/core/model"

type OrderSerializer interface {
	Decode(input []byte) (*model.Order, error)
	DecodeRange(input []byte) ([]*model.Order, error)
	DecodeMap(input []byte) (map[string]interface{}, error)
	DecodeInto(input []byte, target interface{}) error
	Encode(input interface{}) ([]byte, error)
}
