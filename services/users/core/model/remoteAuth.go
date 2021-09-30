package model

type RemoteAuth struct {
	UserID   string             `bson:"user_id"`
	RemoteID string       `bson:"remote_id"`
	Provider UserProvider `bson:"provider"`
}
