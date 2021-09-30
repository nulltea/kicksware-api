package business

import (
	"go.kicksware.com/api/services/users/core/repo"
	"go.kicksware.com/api/services/users/core/service"
)

type interactService struct {
	service   service.UserService
	likesRepo repo.LikesRepository
}

func NewInteractService(userService service.UserService, likesRepo repo.LikesRepository) service.InteractService {
	return &interactService{
		userService,
		likesRepo,
	}
}

func (s *interactService) Like(userID string, entityID string) error {
	return s.likesRepo.AddLike(userID, entityID)
}

func (s *interactService) Unlike(userID string, entityID string) error {
	return s.likesRepo.RemoveLike(userID, entityID)
}
