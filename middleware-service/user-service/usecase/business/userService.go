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

type UserService struct {
	repo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) service.UserService {
	return &UserService{
		userRepo,
	}
}

func (s *UserService) FetchOne(code string) (*model.User, error) {
	return s.repo.FetchOne(code)
}

func (s *UserService) Fetch(codes []string) ([]*model.User, error) {
	return s.repo.Fetch(codes)
}

func (s *UserService) FetchAll() ([]*model.User, error) {
	return s.repo.FetchAll()
}

func (s *UserService) FetchQuery(query meta.RequestQuery) ([]*model.User, error) {
	return s.repo.FetchQuery(query)
}

func (s *UserService) Modify(user *model.User) error {
	return s.repo.Modify(user)
}

func (s *UserService) Replace(user *model.User) error {
	return s.repo.Replace(user)
}

func (s *UserService) Remove(code string) error {
	return s.repo.Remove(code)
}

func (s *UserService) Count(query meta.RequestQuery) (int, error) {
	return s.repo.Count(query)
}

func (s *UserService) CountAll() (int, error) {
	return s.repo.CountAll()
}

func (s *UserService) Register(user *model.User) error {
	if err := validate.Validate(user); err != nil {
		return errs.Wrap(ErrUserInvalid, "service.repo.Register")
	}
	user.UniqueId = xid.New().String()
	user.RegisterDate = time.Now()
	return s.repo.Store(user)
}


