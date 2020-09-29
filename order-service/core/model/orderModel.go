package model

import "time"

type Order struct {
	UniqueID    string      `json:"UniqueID" bson:"unique_id"`
	UserID      string      `json:"UserID" bson:"user_id"`
	ReferenceID string      `json:"ReferenceID" bson:"reference_id"`
	ProductID   string      `json:"ProductID" bson:"product_id"`
	Price       float32     `json:"Price" bson:"price"`
	Status      OrderStatus `json:"Status" bson:"status"`
	SourceURL   string      `json:"SourceURL" bson:"source_url"`
	AddedAt     time.Time   `json:"AddedAt" bson:"added_at"`
	OrderedAt   time.Time   `json:"OrderedAt" bson:"ordered_at"`
}
