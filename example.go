package api

import (
	"github.com/eaigner/hood"
)

// App Schema
type appSchema struct {
	Id      int64
	Name    string
	Guid    string
	SpaceId int64
	StackId int64
}

// App Repository
type AppRepo interface {
	FindByName(name string) (App, error)
	FindAllBySpaceId(id int64) (AppCollection, error)
	Save(Model)
	SaveAll(Collection)
}

type appRepo struct {
	db *hood.Hood
}

func NewAppRepo(hd *hood.Hood) {
	return &appRepo{db: hd}
}

func (repo *appRepo) FindAllBySpaceId(id int64) (apps AppCollection, err error) {
	records := []appSchema{}
	err = repo.db.Where("space_id", "=", id).Limit(1).Find(&records)
	apps = NewAppCollection(records)
	return
}

// App Collection
type AppCollection []App

func NewAppCollection(records []appSchema) AppCollection {
	apps := AppCollection{}
	for record := range records {
		apps = append(coll, NewApp(record))
	}
	return apps
}

// App Model
type App interface {
	Name() string
	SetName(name string)
}

type appModel struct {
	fields appSchema
}

func NewApp(record appSchema) App {
	return &appModel{
		fields: record,
	}
}

func (model *appModel) Name() string {
	return model.fields.Name
}

func (model *appModel) SetName(name string) {
	model.fields.Name = name
}
