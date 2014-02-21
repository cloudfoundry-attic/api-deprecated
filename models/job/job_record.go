package job

import (
	"github.com/cloudfoundry-incubator/api/framework/model"
	"time"
)

type JobRecord struct {
	Id        int64
	Guid      string
	CreatedAt time.Time
	LastError string
	LockedAt  time.Time
}

func (record JobRecord) TableName() string {
	return "delayed_jobs"
}

func NewJobRecord() JobRecord {
	return JobRecord{
		Guid: model.NewGuid(),
	}
}
