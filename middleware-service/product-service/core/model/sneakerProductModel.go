package model

import "time"

type SneakerProduct struct {
	UniqueId       string
	BrandName      string
	ModelName      string
	ModelSKU       string
	ReferenceId    string
	Price          float64
	Type           string
	Size           SneakerSize
	Color          string
	Condition      string
	Description    string
	Owner          string
	Images         map[string][]byte
	ConditionIndex float64
	AddedAt        time.Time
}