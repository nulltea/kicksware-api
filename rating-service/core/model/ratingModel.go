package model

type Rating struct {
	UniqueID string `json:"UniqueID" bson:"unique_id"`
	EntityID string `json:"EntityID" bson:"entity_id"`
	Views    int64  `json:"Views" bson:"views"`
	Orders   int64  `json:"Orders" bson:"orders"`
	Searches int64  `json:"Searches" bson:"searches"`
	Rating   int64  `json:"Rating" bson:"rating"`
}
