package model

import "time"

type User struct {
	UniqueId     string    `json:"uniqueid"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"passwordhash"`
	Confirmed    bool      `json:"confirmed"`
	Admin        bool      `json:"admin"`
	FirstName    string    `json:"firstname"`
	LastName     string    `json:"lastname"`
	RegisterDate time.Time `json:"registerdate"`
	Guest        bool      `json:"-"`
}
