package model

import "time"

type SneakerProduct struct {
	Id        string
	URL       string
	BrandName string
	ModelName string
	Price float32
	Owner     string
	Images    []string
	StateIndex float32
	AddedAt   time.Time
}