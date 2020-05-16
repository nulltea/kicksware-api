package model

type Like struct {
	UserID string `json:"UserID" bson:"user_id"`
	EntityID string `json:"EntityID" bson:"entity_id"`
}
