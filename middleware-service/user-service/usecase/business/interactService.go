package business

import (
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/repo"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/service"
)

type interactService struct {
	service service.UserService
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
