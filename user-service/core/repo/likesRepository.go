package repo

type LikesRepository interface {
	AddLike(userID string, entityID string) error
	RemoveLike(userID string, entityID string) error
}