package business

import (
	"errors"
	errs "github.com/pkg/errors"
	"github.com/rs/xid"
	"gopkg.in/dealancer/validate.v2"
	"time"
	model "user-service/core/model"
	"user-service/core/repo"
	"user-service/core/service"
)

var (
	ErrUserNotFound = errors.New("user Not Found")
	ErrUserInvalid  = errors.New("user Invalid")
)

type UserService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) service.UserService {
	return &UserService{
		userRepo,
	}
}

func (r *UserService) RetrieveOne(code string) (*model.User, error) {
	return r.userRepo.RetrieveOne(code)
}

func (r *UserService) Retrieve(codes []string) ([]*model.User, error) {
	return r.userRepo.Retrieve(codes)
}

func (r *UserService) RetrieveAll() ([]*model.User, error) {
	return r.userRepo.RetrieveAll()
}

func (r *UserService) RetrieveQuery(query interface{}) ([]*model.User, error) {
	return r.userRepo.RetrieveQuery(query)
}

func (r *UserService) Register(user *model.User) error {
	if err := validate.Validate(user); err != nil {
		return errs.Wrap(ErrUserInvalid, "service.userRepo.Store")
	}
	user.UniqueId = xid.New().String()
	user.RegisterDate = time.Now()
	return r.userRepo.Store(user)
}

func (r *UserService) Modify(user *model.User) error {
	return r.userRepo.Modify(user)
}

func (r *UserService) Remove(user *model.User) error {
	return r.userRepo.Remove(user)
}


