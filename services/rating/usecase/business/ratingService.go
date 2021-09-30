package business

import (
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"go.kicksware.com/api/shared/api/rest"
	"go.kicksware.com/api/shared/core"

	"go.kicksware.com/api/services/rating/core/model"
	"go.kicksware.com/api/services/rating/core/repo"
	"go.kicksware.com/api/services/rating/core/service"
	"go.kicksware.com/api/services/rating/env"
)

var (
	ErrRatingNotFound = errors.New("rate Not Found")
	ErrRatingNotValid = errors.New("rate Not Valid")
	uniqueIdFieldName = "unique_id"
)

type ratingService struct {
	repo          repo.RatingRepository
	serviceConfig env.ServiceConfig
	communicator  core.InnerCommunicator
}

func NewRatingService(rateRepo repo.RatingRepository, auth core.AuthService, config env.ServiceConfig) service.RatingService {
	return &ratingService{
		rateRepo,
		config,
		rest.NewCommunicator(auth, config.Common),
	}
}

func (s *ratingService) IncrementViews(entity string) (int64, error) {
	record, err := s.retrieveRecord(entity); if err != nil {
		return 0, err
	}
	record.Views++
	if err := s.repo.Modify(record); err != nil {
		return record.Views, err
	}
	return record.Views, nil
}

func (s *ratingService) IncrementOrders(entity string) (int64, error) {
	record, err := s.retrieveRecord(entity); if err != nil {
		return 0, err
	}
	record.Orders++
	if err := s.repo.Modify(record); err != nil {
		return record.Orders, err
	}
	return record.Orders, nil
}

func (s *ratingService) IncrementSearches(entity string) (int64, error) {
	record, err := s.retrieveRecord(entity); if err != nil {
		return 0, err
	}
	record.Searches++
	if err := s.repo.Modify(record); err != nil {
		return record.Searches, err
	}
	return record.Searches, nil
}

func (s *ratingService) CalculateRating(entity string) (int64, error) {
	record, err := s.retrieveRecord(entity); if err != nil {
		return 0, err
	}
	return calculateRating(record), nil
}

func (s *ratingService) RetrieveRating(entity string) (int64, error) {
	record, err := s.retrieveRecord(entity); if err != nil {
		return 0, err
	}
	return record.Rating, nil
}

func (s *ratingService) UpdateRating(entity string) error {
	panic("implement me")
}

func (s *ratingService) retrieveRecord(entity string) (*model.Rating, error) {
	if record, err := s.repo.FetchOne(entity, nil); record != nil {
		return record, nil
	} else if errors.Cause(err) == ErrRatingNotFound || err == nil {
		record = &model.Rating{
			UniqueID: xid.New().String(),
			EntityID: entity,
		}
		s.repo.StoreOne(record)
		return record, nil

	} else {
		return nil, err
	}
}

func calculateRating(record *model.Rating) int64 {
	return (record.Views + record.Searches * 3 + record.Orders * 5) / 100
}
