package service

import model "go.kicksware.com/api/reference-service/core/model"

type SneakerReferenceSerializer interface {
	Decode(input []byte) (*model.SneakerReference, error)
	DecodeRange(input []byte) ([]*model.SneakerReference, error)
	DecodeMap(input []byte) (map[string]interface{}, error)
	DecodeInto(input []byte, target interface{}) error
	Encode(input interface{}) ([]byte, error)
}
