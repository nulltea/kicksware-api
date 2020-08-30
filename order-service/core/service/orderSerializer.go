package service

import model "github.com/timoth-y/kicksware-platform/middleware-service/order-service/core/model"

type OrderSerializer interface {
	Decode(input []byte) (*model.Order, error)
	DecodeRange(input []byte) ([]*model.Order, error)
	DecodeMap(input []byte) (map[string]interface{}, error)
	DecodeInto(input []byte, target interface{}) error
	Encode(input interface{}) ([]byte, error)
}
