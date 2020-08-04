package service

import (
	"search-service/core/meta"
)

type ProductSyncService interface {
	SyncOne(code string) error
	Sync(codes []string, params *meta.RequestParams) error
	SyncAll(params *meta.RequestParams) error
	SyncQuery(query meta.RequestQuery, params *meta.RequestParams) error
}