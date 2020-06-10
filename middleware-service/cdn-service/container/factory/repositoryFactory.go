package factory

import (
	"github.com/golang/glog"

	"cdn-service/core/repo"
	"cdn-service/env"
	"cdn-service/usecase/storage/files"
	"cdn-service/usecase/storage/mongo"
)

func ProvideRepository(config env.ServiceConfig) repo.ContentRepository {
	switch config.Common.UsedDB {
	case "mongo":
		repo, err := mongo.NewRepository(config.Mongo); if err != nil {
			glog.Fatal(err)
		}
		return repo
	case "files":
		repo, err := files.NewRepository(config.Files); if err != nil {
			glog.Fatal(err)
		}
		return repo
	}
	return nil
}
