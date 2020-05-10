package model

import "time"

type User struct {
	UniqueID     string    `json:"uniqueid"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phonenumber"`
	PasswordHash string    `json:"passwordhash"`
	Confirmed    bool      `json:"confirmed"`
	Admin        bool      `json:"admin"`
	FirstName    string    `json:"firstname"`
	LastName     string    `json:"lastname"`
	RegisterDate time.Time `json:"registerdate"`
	Guest        bool      `json:"-"`
}
