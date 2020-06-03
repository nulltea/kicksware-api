package env

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"gopkg.in/yaml.v2"
)

type ServiceConfig struct {
	Common      CommonConfig    `yaml:"commonConfig"`
	Auth        AuthConfig      `yaml:"authConfig"`
	Mongo       DataStoreConfig `yaml:"mongoConfig"`
	Postgres    DataStoreConfig `yaml:"postgresConfig"`
	Redis       DataStoreConfig `yaml:"redisConfig"`
}

type CommonConfig struct {
	Host               string `yaml:"host"`
	HostName           string `yaml:"hostname"`
	UsedDB             string `yaml:"usedDB"`
	ContentType        string `yaml:"contentType"`
	InnerServiceFormat string `yaml:"innerServiceFormat"`
}

type DataStoreConfig struct {
	URL        string `yaml:"URL"`
	Database   string `yaml:"database"`
	Collection string `yaml:"collection"`
	Login      string `yaml:"login"`
	Password   string `yaml:"password"`
	Timeout    int    `yaml:"timeout"`
}

type AuthConfig struct {
	PublicKeyPath        string `yaml:"publicKeyPath"`
}

func ReadServiceConfig(filename string) (sc ServiceConfig, err error) {
	file, err := ioutil.ReadFile(filename); if err != nil {
		lsR("../")
		log.Fatalln(err)
		return
	}
	err = yaml.Unmarshal(file, &sc); if err != nil {
		return
	}
	return
}

func lsR(searchDir string) ([]string, error) {

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})
	
	if e != nil {
		log.Fatalln(e)
	}

	for _, file := range fileList {
		log.Println(file)
	}

	return fileList, nil
}
