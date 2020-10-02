package service

import "go.kicksware.com/api/service-common/core/meta"

type ReferenceSyncService interface {
	SyncOne(code string) error
	Sync(codes []string, params *meta.RequestParams) error
	SyncAll(params *meta.RequestParams) error
	SyncQuery(query meta.RequestQuery, params *meta.RequestParams) error
}