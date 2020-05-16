package service

type InteractService interface {
	Like(userID string, entityID string) error
	Unlike(userID string, entityID string) error
}
