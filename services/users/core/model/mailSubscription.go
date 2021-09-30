package model

import "time"

type MailSubscription struct {
	Email    string    `bson:"email"`
	UserID   string    `bson:"user_id"`
	Joined   time.Time `bson:"joined"`
}