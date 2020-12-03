package service

import "go.kicksware.com/api/rating-service/core/model"

type RatingSerializer interface {
	Decode(input []byte) (*model.Rating, error)
	DecodeRange(input []byte) ([]*model.Rating, error)
	DecodeMap(input []byte) (map[string]interface{}, error)
	DecodeInto(input []byte, target interface{}) error
	Encode(input interface{}) ([]byte, error)
}
