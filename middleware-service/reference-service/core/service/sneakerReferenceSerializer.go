package service

import model "reference-service/core/model"

type SneakerReferenceSerializer interface {
	DecodeMap(input []byte) (map[string]interface{}, error)
	DecodeOne(input []byte) (*model.SneakerReference, error)
	Decode(input []byte) ([]*model.SneakerReference, error)
	Encode(input interface{}) ([]byte, error)
}
