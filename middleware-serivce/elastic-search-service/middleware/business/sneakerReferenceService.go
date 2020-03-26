package business

import (
	"context"
	"elastic-search-service/core/model"
	"elastic-search-service/core/repo"
	"elastic-search-service/core/service"
	"elastic-search-service/middleware/serializer/json"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
)

const SneakerReferenceIndex = "SneakerReference"

var (
	ErrReferenceNotFound = errors.New("sneaker product Not Found")
)

type searchService struct {
	sneakerProductRepo repo.SneakerReferenceRepository
	client             *elastic.Client
	serializer         service.SneakerReferenceSerializer
}

func NewSneakerReferenceService(sneakerProductRepo repo.SneakerReferenceRepository) service.SneakerReferenceService {
	client, err := elastic.NewClient()
	if err != nil {
		return nil
	}
	if exists, err := client.IndexExists(SneakerReferenceIndex).
		Do(context.Background()); err != nil || !exists {
		_, err = client.CreateIndex(SneakerReferenceIndex).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
	return &searchService{
		sneakerProductRepo,
		client,
		json.NewSerializer(),
	}
}

func (s *searchService) SyncOne(code string) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ref, err := s.sneakerProductRepo.FetchOne(code)
	if err != nil {
		return errors.Wrap(err, "service.SyncOne")
	}

	body, err := s.serializer.Encode(ref)
	if err != nil {
		return errors.Wrap(err, "service.SyncOne")
	}

	_, err = s.client.Index().
		Index(SneakerReferenceIndex).
		Type("doc").
		Id(ref.UniqueId).
		BodyJson(body).
		Refresh("wait_for").
		Do(ctx)
	if err != nil {
		return errors.Wrap(err, "service.SyncOne")
	}
	return nil
}

func (s *searchService) SyncQuery(query interface{}) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	refs := make([]*model.SneakerReference, 0)

	if query == nil {
		refs, err = s.sneakerProductRepo.FetchAll()
		if err != nil {
			return
		}
	}

	if codes := query.([]string); codes != nil {
		refs, err = s.sneakerProductRepo.Fetch(codes)
		if err != nil {
			return
		}
	}

	refs, err = s.sneakerProductRepo.FetchQuery(query)
	if err != nil {
		return
	}

	for _, ref := range refs {
		body, err := s.serializer.Encode(refs)
		if err != nil {
			continue
		}

		_, err = s.client.Index().
			Index(SneakerReferenceIndex).
			Type("doc").
			Id(ref.UniqueId).
			BodyJson(body).
			Refresh("wait_for").
			Do(ctx)
		if err != nil {
			continue
		}
	}
	return nil
}

func (s *searchService) Sync(codes []string) (err error) {
	if err = s.SyncQuery(codes); err != nil {
		return
	}
	return
}

func (s *searchService) SyncAll() (err error) {
	if err = s.SyncQuery(nil); err != nil {
		return
	}
	return
}
