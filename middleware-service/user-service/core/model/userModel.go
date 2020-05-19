package model

import "time"

type User struct {
	UniqueID     string       `json:"UniqueID" bson:"unique_id"`
	Username     string       `json:"Username" bson:"username"`
	UsernameN    string       `json:"UsernameN" bson:"username_n"`
	Email        string       `json:"Email" bson:"email"`
	EmailN       string       `json:"EmailN" bson:"email_n"`
	PasswordHash string       `json:"PasswordHash" bson:"password_hash"`
	FirstName    string       `json:"FirstName" bson:"first_name"`
	LastName     string       `json:"LastName" bson:"last_name"`
	PhoneNumber  string       `json:"PhoneNumber" bson:"phone_number"`
	Avatar       string       `json:"Avatar" bson:"avatar"`
	Location     string       `json:"Location" bson:"location"`
	PaymentInfo  PaymentInfo  `json:"PaymentInfo" bson:"payment_info"`
	Liked        []string     `json:"Liked,omitempty" bson:"liked,omitempty"`
	Settings     UserSettings `json:"Settings,omitempty" bson:"settings,omitempty"`
	Confirmed    bool         `json:"Confirmed" bson:"confirmed"`
	Role         UserRole     `json:"Role" bson:"role"`
	RegisterDate time.Time    `json:"RegisterDate" bson:"register_date"`
}
