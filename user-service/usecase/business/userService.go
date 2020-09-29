package business

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	errs "github.com/pkg/errors"
	"github.com/rs/xid"
	"gopkg.in/dealancer/validate.v2"

	"github.com/timoth-y/kicksware-api/service-common/core/meta"

	"github.com/timoth-y/kicksware-api/user-service/core/model"
	"github.com/timoth-y/kicksware-api/user-service/core/repo"
	"github.com/timoth-y/kicksware-api/user-service/core/service"
)

var (
	ErrUserNotFound = errors.New("user Not Found")
	ErrUserInvalid  = errors.New("user Invalid")
	ErrEmailInvalid  = errors.New("user email could not be empty")
)

type userService struct {
	repo       repo.UserRepository
	remoteRepo repo.RemoteRepository
}

func NewUserService(userRepo repo.UserRepository, remoteRepo repo.RemoteRepository) service.UserService {
	return &userService{
		userRepo,
		remoteRepo,
	}
}

func (s *userService) FetchOne(code string) (*model.User, error) {
	return s.repo.FetchOne(code)
}

func (s *userService) FetchByEmail(email string) (*model.User, error) {
	query := meta.RequestQuery{"email": email}
	users, err := s.repo.FetchQuery(query, nil); if err != nil || len(users) == 0 {
		return nil, err
	}
	return users[0], nil
}

func (s *userService) FetchByUsername(username string) (*model.User, error) {
	query := meta.RequestQuery{"username": username}
	users, err := s.repo.FetchQuery(query, nil); if err != nil || len(users) == 0 {
		return nil, err
	}
	return users[0], nil
}

func (s *userService) Fetch(codes []string, params *meta.RequestParams) ([]*model.User, error) {
	return s.repo.Fetch(codes, params)
}

func (s *userService) FetchAll(params *meta.RequestParams) ([]*model.User, error) {
	return s.repo.FetchAll(params)
}

func (s *userService) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) ([]*model.User, error) {
	return s.repo.FetchQuery(query, params)
}

func (s *userService) FetchRemote(remoteID string, provider model.UserProvider) (*model.User, error) {
	userID, err := s.remoteRepo.Track(remoteID, provider); if err != nil {
		return nil, err
	}
	return s.FetchOne(userID)
}

func (s *userService) GenerateUsername(user *model.User, save bool) (string, error) {
	if len(user.Email) == 0 {
		return "", ErrEmailInvalid
	}
	username := strings.Split(user.Email, "@")[0]
	if another, _ := s.FetchByUsername(username); another != nil {
		baseUsername := username
		for another != nil {
			rand.Seed(user.RegisterDate.Unix())
			username = fmt.Sprintf("%v_%v", baseUsername, strconv.Itoa(rand.Int())[:3])
			another, _ = s.FetchByUsername(username)
		}
	}
	user.Username = username
	if save {
		if err := s.Modify(user); err != nil {
			return "", err
		}
	}
	return username, nil;
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

func (s *userService) Count(query meta.RequestQuery, params *meta.RequestParams) (int, error) {
	return s.repo.Count(query, params)
}

func (s *userService) CountAll() (int, error) {
	return s.repo.CountAll()
}

func (s *userService) Register(user *model.User) error {
	if err := validate.Validate(user); err != nil {
		return errs.Wrap(ErrUserInvalid, "service.repo.Register")
	}
	user.RegisterDate = time.Now()
	if len(user.UniqueID) < 8 {
		user.UniqueID = xid.NewWithTime(user.RegisterDate).String()
	}
	if len(user.Username) == 0 {
		s.GenerateUsername(user, false)
	}

	if err := s.repo.Store(user); err != nil {
		return err
	}

	if user.ConnectedProviders != nil && len(user.ConnectedProviders) != 0 {
		if err := s.remoteRepo.Sync(user.UniqueID, user.ConnectedProviders); err != nil {
			return err
		}
	}
	return nil
}

func (s *userService) ConnectProvider(userID string, remoteID string, provider model.UserProvider) error {
	return s.remoteRepo.Connect(userID, remoteID, provider)
}


