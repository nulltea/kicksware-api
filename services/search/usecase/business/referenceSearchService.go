package business

import (
	"context"
	"log"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"

	"go.kicksware.com/api/services/references/core/model"

	"go.kicksware.com/api/shared/core/meta"

	"go.kicksware.com/api/services/search/core/service"
	"go.kicksware.com/api/services/search/env"
	"go.kicksware.com/api/services/search/usecase/serializer/json"
)

var (
	ErrReferenceNotFound = errors.New("sneaker reference Not Found")
	ErrReferenceNotValid = errors.New("sneaker reference Not Valid")
)

type referenceSearchService struct {
	client     *elastic.Client
	serializer service.SneakerSearchSerializer
	index      string
	params     env.SearchConfig
}

func NewReferenceSearchService(config env.ElasticConfig, params env.SearchConfig) service.ReferenceSearchService {
	client, err := initElasticSearchClient(config); if err != nil {
		log.Fatalln(err)
		return nil
	}
	return &referenceSearchService{
		client,
		json.NewSerializer(),
		config.Index,
		params,
	}
}

func initElasticSearchClient(config env.ElasticConfig) (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(config.URL),
		elastic.SetMaxRetries(10), // TODO retry implementation
		elastic.SetHealthcheckTimeoutStartup(time.Duration(config.StartupDelay) * time.Second),
		elastic.SetSniff(config.Sniffing),
	); if err != nil {
		return nil, err
	}
	if exists, err := client.IndexExists(config.Index).
		Do(context.Background()); err != nil || !exists {
		_, err = client.CreateIndex(config.Index).Do(context.Background()); if err != nil {
			return nil, err
		}
	}
	return client, nil
}

func (s *referenceSearchService) Search(query string, params *meta.RequestParams) ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	matchQuery := elastic.NewMultiMatchQuery(query, s.params.Fields...).
		Type(s.params.Type).
		Slop(s.params.Slop)

	if params == nil {
		params = &meta.RequestParams{}
	}
	sortBy := params.SortBy(); if len(sortBy) == 0 {
		sortBy = "Price"
	}

	results, err := s.client.Search().
		Index(s.index).
		Query(matchQuery).
		Sort(sortBy, params.SortDirection() == "asc").
		From(params.Offset()).
		Size(params.Limit()).
		Do(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "service.Search")
	}

	refs := make([]*model.SneakerReference, 0)
	if results.TotalHits() < 0 {
		return refs, nil
	}
	for _, hit := range results.Hits.Hits {
		if ref, err := s.serializer.DecodeReference(hit.Source); err == nil {
			refs = append(refs, ref)
		}
	}
	return refs, nil
}

func (s *referenceSearchService) SearchBy(field, value string, params *meta.RequestParams) ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	matchQuery := elastic.NewMatchQuery(field, value)
	if params == nil {
		params = &meta.RequestParams{}
	}
	sortBy := params.SortBy(); if len(sortBy) != 0 {
		sortBy = "Price"
	}

	results, err := s.client.Search().
		Index(s.index).
		Query(matchQuery).
		Sort(params.SortBy(), params.SortBy() == "asc").
		From(params.Offset()).
		Size(params.Limit()).
		Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "service.searchBy")
	}
	log.Printf(
		"Query on %v with %v took %d milliseconds\n found total %v",
		field, value, results.TotalHits(), results.TookInMillis,
	)

	refs := make([]*model.SneakerReference, 0)
	if results.TotalHits() < 0 {
		return refs, nil
	}
	for _, hit := range results.Hits.Hits {
		if ref, err := s.serializer.DecodeReference(hit.Source); err == nil {
			refs = append(refs, ref)
		}
	}
	return refs, nil
}

func (s *referenceSearchService) SearchSKU(sku string, params *meta.RequestParams) (refs []*model.SneakerReference, err error) {
	if refs, err = s.SearchBy("ManufactureSku", sku, params); err != nil {
		return nil, errors.Wrap(err, "service.SearchSKU")
	}
	return
}

func (s *referenceSearchService) SearchBrand(brand string, params *meta.RequestParams) (refs []*model.SneakerReference, err error) {
	if refs, err = s.SearchBy("BrandName", brand, params); err != nil {
		return nil, errors.Wrap(err, "service.SearchBrand")
	}
	return
}

func (s *referenceSearchService) SearchModel(model string, params *meta.RequestParams) (refs []*model.SneakerReference, err error) {
	if refs, err = s.SearchBy("ModelName", model, params); err != nil {
		return nil, errors.Wrap(err, "service.SearchModel")
	}
	return
}
