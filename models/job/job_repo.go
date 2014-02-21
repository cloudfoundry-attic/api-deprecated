package job

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type JobRepo interface {
	FindByGuid(guid string) (job JobModel, found bool)
	Save(JobModel) error
}

type jobRepo struct {
	db gorm.DB
}

func NewRepo(db gorm.DB) JobRepo {
	return &jobRepo{
		db: db,
	}
}

func (repo *jobRepo) FindByGuid(guid string) (job JobModel, found bool) {
	jobRecord := new(JobRecord)
	dbCon := repo.db.Where("guid = ?", guid)
	dbCon.First(jobRecord)
	if jobRecord.Id == 0 {
		return
	}
	job = NewModelFromRecord(*jobRecord)
	found = true
	return
}

func (repo *jobRepo) Save(m JobModel) (err error) {
	recorder, ok := m.(Recorder)
	if !ok {
		err = errors.New(fmt.Sprintf("Model %T must implement Recorder to save.", m))
	}

	// TODO: UTC timestamps need to be fixed in GORM
	// https://github.com/jinzhu/gorm/pull/67
	record := recorder.Record().(*JobRecord)
	record.LockedAt = record.LockedAt.UTC()
	repo.db.Save(record)
	recorder.SetRecord(record)
	return
}
