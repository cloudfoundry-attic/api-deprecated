package models

import (
	"github.com/cloudfoundry-incubator/api/models/job"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func NewExports(db gorm.DB) map[interface{}]interface{} {
	return map[interface{}]interface{}{
		(*job.JobRepo)(nil): job.NewRepo(db),
	}
}
