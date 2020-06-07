package model

import "time"

type Order struct {
	UniqueID  string      `json:"UniqueID" bson:"unique_id"`
	UserID    string      `json:"UserID" bson:"user_id"`
	ProductID string      `json:"ProductID" bson:"product_id"`
	Price     float64     `json:"Price" bson:"price"`
	Status    OrderStatus `json:"Status" bson:"status"`
	OrderedAt time.Time   `json:"OrderedAt" bson:"ordered_at"`
}
