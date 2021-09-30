package model

type UserSettings struct {
	Theme string `json:"Theme" bson:"theme"`
	LayoutView string `json:"LayoutView" bson:"layout_view"`
}
