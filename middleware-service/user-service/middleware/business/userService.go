package business

import (
	"errors"
	errs "github.com/pkg/errors"
	"github.com/rs/xid"
	"gopkg.in/dealancer/validate.v2"
	"time"
	"user-service/core/meta"
	"user-service/core/model"
	"user-service/core/repo"
	"user-service/core/service"
)

var (
	ErrUserNotFound = errors.New("user Not Found")
	ErrUserInvalid  = errors.New("user Invalid")
)

type userService struct {
	repo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) service.UserService {
	return &userService{
		userRepo,
	}
}

func (s *userService) FetchOne(code string) (*model.User, error) {
	return s.repo.FetchOne(code)
}

func (s *userService) Fetch(codes []string) ([]*model.User, error) {
	return s.repo.Fetch(codes)
}

func (s *userService) FetchAll() ([]*model.User, error) {
	return s.repo.FetchAll()
}

func (s *userService) FetchQuery(query meta.RequestQuery) ([]*model.User, error) {
	return s.repo.FetchQuery(query)
}

func (s *userService) Modify(user *model.User) error {
	return s.repo.Modify(user)
}

func (s *userService) Replace(user *model.User) error {
	return s.repo.Replace(user)
}

func (s *userService) Remove(code string) error {
	return s.repo.Remove(code)
}

func (s *userService) Count(query meta.RequestQuery) (int, error) {
	return s.repo.Count(query)
}

func (s *userService) CountAll() (int, error) {
	return s.repo.CountAll()
}

func (s *userService) Register(user *model.User) error {
	if err := validate.Validate(user); err != nil {
		return errs.Wrap(ErrUserInvalid, "service.repo.Register")
	}
	user.UniqueId = xid.New().String()
	user.RegisterDate = time.Now()
	return s.repo.Store(user)
}


