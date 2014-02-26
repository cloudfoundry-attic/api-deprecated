package job

import (
	"github.com/cloudfoundry-incubator/api/framework/models"
	"time"
)

type Record struct {
	Id        int64
	Guid      string
	CreatedAt time.Time
	LastError string
	LockedAt  time.Time
}

func (record Record) TableName() string {
	return "delayed_jobs"
}

func NewRecord() Record {
	return Record{
		Guid: models.NewGuid(),
	}
}
