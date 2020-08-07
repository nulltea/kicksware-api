package disk

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/glog"

	"github.com/timoth-y/kicksware-platform/middleware-service/cdn-service/core/repo"
	"github.com/timoth-y/kicksware-platform/middleware-service/cdn-service/env"
)

type repository struct {
	storagePath string
}

func NewRepository(config env.DataStoreConfig) (repo.ContentRepository, error) {
	path := config.URL
	if _, err := os.Stat(path); err != nil && !os.IsExist(err) {
		glog.Fatalln(err)
		return nil, err
	}

	return &repository{path}, nil
}


func (r *repository) Download(from string, filename string) ([]byte, error) {
	file, err := ioutil.ReadFile(r.filenameOf(from, filename))
	if err != nil {
		glog.Errorln(err)
		return nil, err
	}

	return file, nil
}

func (r *repository) Upload(to string, filename string, content []byte) (string, error) {
	err := ioutil.WriteFile(r.filenameOf(to, filename), content, 0600); if err != nil {
		glog.Errorln(err)
		return "", err
	}
	return filename, nil
}

func (r *repository) filenameOf(collection, filename string) string {
	return fmt.Sprintf("%v/%v/%v", r.storagePath, collection, filename)
}