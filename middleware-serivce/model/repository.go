package model

type SneakerProductRepository interface {
	Find(code string) (*SneakerProduct, error)
	Store(sneakerProduct *SneakerProduct) error
}

