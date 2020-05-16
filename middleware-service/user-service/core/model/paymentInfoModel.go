package model

type PaymentInfo struct {
	CardNumber  string      `json:"CardNumber" bson:"card_number"`
	Expires     string      `json:"Expires" bson:"expires"`
	CVV         string      `json:"CVV" bson:"cvv"`
	BillingInfo AddressInfo `json:"BillingInfo" bson:"billing_info"`
}
