package model

import "time"

type SneakerSize struct {
	Europe        float64
	UnitedStates  float64
	UnitedKingdom float64
	Centimeters   float64
}

type SneakerProduct struct {
	UniqueId       string
	BrandName      string
	ModelName      string
	Price          float64
	Type           string
	Size           SneakerSize
	Color          string
	Condition      string
	Description    string
	Owner          string
	Images         []string
	ConditionIndex float64
	AddedAt        time.Time
}