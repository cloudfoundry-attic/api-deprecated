package app

import (
	"github.com/cloudfoundry-incubator/api/framework/models"
)

type Record struct {
	Id           int64
	Guid         string
	PackageHash  string
	PackageState string
}

func (record Record) TableName() string {
	return "apps"
}

func NewRecord() Record {
	return Record{
		Guid: models.NewGuid(),
	}
}
