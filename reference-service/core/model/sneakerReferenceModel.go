package model

import "time"

type SneakerReference struct {
	UniqueId       string
	ManufactureSku string
	BrandName      string
	Brand          SneakerBrand
	ModelName      string
	Model          SneakerModel
	BaseModelName  string
	BaseModel      SneakerModel
	Description    string
	Color          string
	Gender         string
	Nickname       string
	Materials      []string
	Categories     []string
	ReleaseDate    time.Time
	ReleaseDateStr string `bson:"release_strdate"`
	Price          float64
	ImageLink      string
	ImageLinks     []string
	StadiumUrl     string
	Likes          int
	Liked          bool
	AddedDate      time.Time `bson:"added_date"`
}