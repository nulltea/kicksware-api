package model

type SneakerProductSerializer interface {
	Decode(input []byte) (*SneakerProduct, error)
	Encode(input *SneakerProduct) ([]byte, error)
}
