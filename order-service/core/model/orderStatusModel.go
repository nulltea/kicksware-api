package model

type OrderStatus string

var (
	Draft      OrderStatus = "draft"
	Processing OrderStatus = "processing"
	OnHold     OrderStatus = "on_hold"
	Delivering OrderStatus = "delivering"
	Complete   OrderStatus = "complete"
)
