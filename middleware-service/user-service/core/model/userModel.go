package model

import "time"

type User struct {
	UniqueID     string    `json:"uniqueid"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phonenumber"`
	PasswordHash string    `json:"passwordhash"`
	Confirmed    bool      `json:"confirmed"`
	Role         UserRole  `json:"role"`
	FirstName    string    `json:"firstname"`
	LastName     string    `json:"lastname"`
	RegisterDate time.Time `json:"registerdate"`
}
