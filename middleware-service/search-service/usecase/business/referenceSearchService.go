package business

import (
	"context"
	"log"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"

	"search-service/core/model"
	"search-service/core/service"
	"search-service/env"
	"search-service/usecase/serializer/json"
)

var (
	ErrReferenceNotFound = errors.New("sneaker reference Not Found")
	ErrReferenceNotValid = errors.New("sneaker reference Not Valid")
)

type referenceSearchService struct {
	client             *elastic.Client
	serializer         service.SneakerSearchSerializer
	index              string
}

func NewReferenceSearchService(config env.ElasticConfig) service.ReferenceSearchService {
	client, err := initElasticSearchClient(config); if err != nil {
		log.Fatalln(err)
		return nil
	}
	return &referenceSearchService{
		client,
		json.NewSerializer(),
		config.Index,
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

func (s *referenceSearchService) Search(query string) ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	matchQuery := elastic.NewMultiMatchQuery(query)
	results, err := s.client.Search().
		Index(s.index).
		Query(matchQuery).
		Sort("Price", false).
		From(0).
		Size(10).
		Do(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "service.Search")
	}
	log.Printf(		"Query on %v took %d milliseconds\n found total %v",
		query, results.TotalHits(), results.TookInMillis,
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

func (s *referenceSearchService) SearchBy(field, value string) ([]*model.SneakerReference, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	matchQuery := elastic.NewMatchQuery(field, value)
	results, err := s.client.Search().
		Index(s.index).
		Query(matchQuery).
		Sort("Price", false).
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

func (s *referenceSearchService) SearchSKU(sku string) (refs []*model.SneakerReference, err error) {
	if refs, err = s.SearchBy("ManufactureSku", sku); err != nil {
		return nil, errors.Wrap(err, "service.SearchSKU")
	}
	return
}

func (s *referenceSearchService) SearchBrand(brand string) (refs []*model.SneakerReference, err error) {
	if refs, err = s.SearchBy("BrandName", brand); err != nil {
		return nil, errors.Wrap(err, "service.SearchBrand")
	}
	return
}

func (s *referenceSearchService) SearchModel(model string) (refs []*model.SneakerReference, err error) {
	if refs, err = s.SearchBy("ModelName", model); err != nil {
		return nil, errors.Wrap(err, "service.SearchModel")
	}
	return
}
