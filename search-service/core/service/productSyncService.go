package service

import (
	"github.com/timoth-y/kicksware-api/service-common/core/meta"
)

type ProductSyncService interface {
	SyncOne(code string) error
	Sync(codes []string, params *meta.RequestParams) error
	SyncAll(params *meta.RequestParams) error
	SyncQuery(query meta.RequestQuery, params *meta.RequestParams) error
}