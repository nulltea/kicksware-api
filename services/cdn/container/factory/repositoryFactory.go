package factory

import (
	"github.com/golang/glog"

	"go.kicksware.com/api/services/cdn/core/repo"
	"go.kicksware.com/api/services/cdn/env"
	"go.kicksware.com/api/services/cdn/usecase/storage/disk"
	"go.kicksware.com/api/services/cdn/usecase/storage/mongo"
)

func ProvideRepository(config env.ServiceConfig) repo.ContentRepository {
	switch config.Common.UsedDB {
	case "mongo":
		repo, err := mongo.NewRepository(config.Mongo); if err != nil {
			glog.Fatal(err)
		}
		return repo
	case "disk":
		repo, err := disk.NewRepository(config.Files); if err != nil {
			glog.Fatal(err)
		}
		return repo
	}
	return nil
}
