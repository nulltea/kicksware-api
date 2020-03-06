package model

type SneakerProductService interface {
	Find(code string) (*SneakerProduct, error)
	Store(sneakerProduct *SneakerProduct) error
}