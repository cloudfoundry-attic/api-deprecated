package config

import (
	"github.com/fraenkel/candiedyaml"
	"io/ioutil"
)

type Config struct {
	DefaultBackendURL string `yaml:"default_backend_url"`
	Port              int
	DB                DbConfig
	AppPackages       BlobstoreConfig `yaml:"app_packages"`
}

type DbConfig struct {
	URI string `yaml:"database"`
}

type BlobstoreConfig struct {
	Filepath string
}

func New(configBytes []byte) (c Config, err error) {
	err = candiedyaml.Unmarshal(configBytes, &c)
	return
}

func NewFromFile(filePath string) (c Config, err error) {
	configBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	return New(configBytes)
}
