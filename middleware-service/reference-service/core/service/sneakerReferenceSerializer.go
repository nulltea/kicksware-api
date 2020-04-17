package service

import model "reference-service/core/model"

type SneakerReferenceSerializer interface {
	DecodeOne(input []byte) (*model.SneakerReference, error)
	Decode(input []byte) ([]*model.SneakerReference, error)
	Encode(input interface{}) ([]byte, error)
}
