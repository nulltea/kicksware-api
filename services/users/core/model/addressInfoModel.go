package model

type AddressInfo struct {
	Country    string `json:"Country" bson:"country"`
	City       string `json:"City" bson:"city"`
	Address    string `json:"Address" bson:"address"`
	Address2   string `json:"Address2" bson:"address2"`
	Region     string `json:"Region" bson:"region"`
	PostalCode string `json:"PostalCode" bson:"postalCode"`
}
