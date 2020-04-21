package model

import "time"

type SneakerReference struct {
	UniqueId         string
	ManufactureSku   string
	BrandName        string
	ModelName        string
	Description      string
	Color            string
	Gender           string
	Nickname         string
	Price            float64
	Released         time.Time
	ImageLink        string
	StadiumUrl       string
}