package business

import (
	"context"
	"log"

	"github.com/timoth-y/kicksware-platform/middleware-service/reference-service/core/model"

	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/meta"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/pipe"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/core/service"
	"github.com/timoth-y/kicksware-platform/middleware-service/search-service/env"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
)

var(
	ErrInvalidQuery = errors.New("searchReferenceService: query object must be either a string, []string or a meta.RequestQuery")
)

type referenceSyncService struct {
	pipe   pipe.SneakerReferencePipe
	client *elastic.Client
	index  string
}

func NewReferenceSyncService(pipe pipe.SneakerReferencePipe, config env.ElasticConfig) service.ReferenceSyncService {
	client, err := initElasticSearchClient(config); if err != nil {
		log.Fatalln(err)
		return nil
	}
	return &referenceSyncService{
		pipe,
		client,
		config.Index,
	}
}

func (s *referenceSyncService) SyncOne(code string) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ref, err := s.pipe.FetchOne(code)
	if err != nil {
		return errors.Wrap(err, "service.SyncOne")
	}

	_, err = s.client.Index().
		Index(s.index).
		Id(ref.UniqueId).
		BodyJson(ref).
		Refresh("wait_for").
		Do(ctx)
	if err != nil {
		return errors.Wrap(err, "service.SyncOne")
	}
	return nil
}

func (s *referenceSyncService) Sync(codes []string, params *meta.RequestParams) (err error) {
	if err = s.sync(codes, params); err != nil {
		return
	}
	return
}

func (s *referenceSyncService) SyncAll(params *meta.RequestParams) (err error) {
	if err = s.sync(nil, params); err != nil {
		return
	}
	return
}

func (s *referenceSyncService) SyncQuery(query meta.RequestQuery, params *meta.RequestParams) (err error) {
	if err = s.sync(query, params); err != nil {
		return
	}
	return
}

func (s *referenceSyncService) sync(query interface{}, params *meta.RequestParams) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	refs := make([]*model.SneakerReference, 0)

	switch qv := query.(type) {
	case nil:
		refs, err = s.pipe.FetchAll(params); if err != nil {
		return
	}
	case []string:
		refs, err = s.pipe.Fetch(qv, params); if err != nil {
		return
	}
	case meta.RequestQuery:
		refs, err = s.pipe.FetchQuery(qv, params); if err != nil {
		return
	}
	default:
		return errors.Wrap(ErrInvalidQuery, "service.SyncQuery")
	}
	bulk := s.client.Bulk()
	for _, ref := range refs {
		bulk.Add(
			elastic.NewBulkIndexRequest().
				Index(s.index).
				Id(ref.UniqueId).
				Doc(ref),
		)
	}
	if _, err := bulk.Do(ctx); err != nil {
		return errors.Wrap(err, "service.SyncQuery")
	}
	return nil
}