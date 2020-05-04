package model

import "time"

type User struct {
	UniqueId     string
	Username     string
	PasswordHash string
	Confirmed    bool
	Admin        bool
	FirstName    string
	LastName     string
	RegisterDate time.Time
}
