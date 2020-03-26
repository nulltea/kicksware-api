package service

type SneakerReferenceService interface {
	SyncOne(code string) error
	Sync(codes []string) error
	SyncAll() error
	SyncQuery(query interface{}) error
}