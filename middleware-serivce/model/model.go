package model

import "time"

type SneakerProduct struct {
	Id        string
	URL       string
	BrandName string
	ModelName string
	Owner     string
	Images    []string
	StateIndex float32
	AddedAt   time.Time
}