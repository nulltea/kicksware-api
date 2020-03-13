package model

import "time"

type User struct {
	UniqueId     string
	UserName     string
	PasswordHash string
	FirstName    string
	LastName     string
	RegisterDate time.Time
}
