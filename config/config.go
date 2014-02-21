package config

import (
	"io/ioutil"
	"launchpad.net/goyaml"
)

type Config struct {
	DefaultBackendURL string `yaml:"default_backend_url"`
	Port              int
	DB                DbConfig
}

type DbConfig struct {
	URI string `yaml:"database"`
}

func New(configBytes []byte) (c Config, err error) {
	err = goyaml.Unmarshal(configBytes, &c)
	return
}

func NewFromFile(filePath string) (c Config, err error) {
	configBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	return New(configBytes)
}
