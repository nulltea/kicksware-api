package business

import (
	"context"
	"elastic-search-service/core/model"
	"elastic-search-service/core/repo"
	"elastic-search-service/core/service"
	"elastic-search-service/middleware/serializer/json"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"log"
	"time"
)

var (
	ErrReferenceNotFound = errors.New("sneaker reference Not Found")
	ErrReferenceNotValid = errors.New("sneaker reference Not Valid")
)

type searchService struct {
	sneakerProductRepo repo.SneakerReferenceRepository
	client             *elastic.Client
	serializer         service.SneakerReferenceSerializer
	index              string
}

func NewSneakerReferenceService(sneakerProductRepo repo.SneakerReferenceRepository, elasticURL, elasticIndex string) service.SneakerReferenceService {
	client, err := elastic.NewClient(
		elastic.SetURL(elasticURL),
		elastic.SetMaxRetries(10),
		elastic.SetHealthcheckTimeoutStartup(60*time.Second),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Println(err)
		return nil
	}
	if exists, err := client.IndexExists(elasticIndex).
		Do(context.Background()); err != nil || !exists {
		_, err = client.CreateIndex(elasticIndex).Do(context.Background())
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
	return &searchService{
		sneakerProductRepo,
		client,
		json.NewSerializer(),
		elasticIndex,
	}
}

func (s *searchService) SyncOne(code string) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ref, err := s.sneakerProductRepo.FetchOne(code)
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

func (s *searchService) SyncQuery(query interface{}) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	refs := make([]*model.SneakerReference, 0)

	switch qv := query.(type) {
	case nil:
		refs, err = s.sneakerProductRepo.FetchAll()
		if err != nil {
			return
		}
	case []string:
		refs, err = s.sneakerProductRepo.Fetch(qv)
		if err != nil {
			return
		}
	default:
		refs, err = s.sneakerProductRepo.FetchQuery(qv)
		if err != nil {
			return
		}
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

func (s *searchService) Search(query string) ([]*model.SneakerReference, error) {
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
		if ref, err := s.serializer.Decode(hit.Source); err == nil {
			refs = append(refs, ref)
		}
	}
	return refs, nil
}

func (s *searchService) SearchBy(field, value string) ([]*model.SneakerReference, error) {
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
		if ref, err := s.serializer.Decode(hit.Source); err == nil {
			refs = append(refs, ref)
		}
	}
	return refs, nil
}

func (s *searchService) SearchSKU(sku string) (refs []*model.SneakerReference, err error) {
	if refs, err = s.SearchBy("ManufactureSku", sku); err != nil {
		return nil, errors.Wrap(err, "service.SearchSKU")
	}
	return
}

func (s *searchService) SearchBrand(brand string) (refs []*model.SneakerReference, err error) {
	if refs, err = s.SearchBy("BrandName", brand); err != nil {
		return nil, errors.Wrap(err, "service.SearchBrand")
	}
	return
}

func (s *searchService) SearchModel(model string) (refs []*model.SneakerReference, err error) {
	if refs, err = s.SearchBy("ModelName", model); err != nil {
		return nil, errors.Wrap(err, "service.SearchModel")
	}
	return
}
