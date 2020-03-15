package business

import (
	"errors"
	errs "github.com/pkg/errors"
	"github.com/rs/xid"
	"gopkg.in/dealancer/validate.v2"
	"time"
	"user-service/core/model"
	"user-service/core/repo"
	"user-service/core/service"
)

var (
	ErrUserNotFound = errors.New("user Not Found")
	ErrUserInvalid  = errors.New("user Invalid")
)

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) service.UserService {
	return &userService{
		userRepo,
	}
}

func (r *userService) FetchOne(code string) (*model.User, error) {
	return r.userRepo.FetchOne(code)
}

func (r *userService) Fetch(codes []string) ([]*model.User, error) {
	return r.userRepo.Fetch(codes)
}

func (r *userService) FetchAll() ([]*model.User, error) {
	return r.userRepo.FetchAll()
}

func (r *userService) FetchQuery(query interface{}) ([]*model.User, error) {
	return r.userRepo.FetchQuery(query)
}

func (r *userService) Register(user *model.User) error {
	if err := validate.Validate(user); err != nil {
		return errs.Wrap(ErrUserInvalid, "service.userRepo.Register")
	}
	user.UniqueId = xid.New().String()
	user.RegisterDate = time.Now()
	return r.userRepo.Store(user)
}

func (r *userService) Modify(user *model.User) error {
	return r.userRepo.Modify(user)
}

func (r *userService) Replace(user *model.User) error {
	return r.userRepo.Replace(user)
}

func (r *userService) Remove(code string) error {
	return r.userRepo.Remove(code)
}

func (r *userService) RemoveObj(user *model.User) error {
	return r.userRepo.RemoveObj(user)
}


