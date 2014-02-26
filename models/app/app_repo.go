package app

import (
	"errors"
	"fmt"
	"github.com/cloudfoundry-incubator/api/framework/models"

	"github.com/jinzhu/gorm"
)

type Repo interface {
	FindByGuid(guid string) (app Model, found bool)
	Save(Model) error
}

type repo struct {
	db gorm.DB
}

func NewRepo(db gorm.DB) Repo {
	return &repo{
		db: db,
	}
}

func (r *repo) FindByGuid(guid string) (m Model, found bool) {
	data := new(Record)
	dbCon := r.db.Where("guid = ?", guid)
	dbCon.First(data)
	if data.Id == 0 {
		return
	}
	m = NewModelFromRecord(*data)
	found = true
	return
}

func (r *repo) Save(m Model) (err error) {
	recorder, ok := m.(models.Recorder)
	if !ok {
		err = errors.New(fmt.Sprintf("Model %T must implement Recorder to save.", m))
		return
	}

	record := recorder.Record()
	r.db.Save(record)
	recorder.SetRecord(record)
	return
}
