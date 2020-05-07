package model

import "time"

type SneakerReference struct {
	UniqueId       string
	ManufactureSku string
	BrandName      string
	ModelName      string
	BaseModelName  string
	Description    string
	Color          string
	Gender         string
	Nickname       string
	Materials      []string
	Categories     []string
	ReleaseDate    time.Time
	Price          float64
	ImageLink      string
	ImageLinks     []string
	StadiumUrl     string
}