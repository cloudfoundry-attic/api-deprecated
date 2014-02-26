package fake_app

import (
	"github.com/cloudfoundry-incubator/api/models/app"
)

type FakeRepo struct {
	FindByGuidInputs struct {
		Guid string
	}
	FindByGuidOutputs struct {
		Model app.Model
		Found bool
	}
	SaveInputs struct {
		Model app.Model
	}
	SaveOutputs struct {
		Err error
	}
}

func NewFakeRepo() *FakeRepo {
	return new(FakeRepo)
}

func (r *FakeRepo) FindByGuid(guid string) (m app.Model, found bool) {
	r.FindByGuidInputs.Guid = guid
	m = r.FindByGuidOutputs.Model
	found = r.FindByGuidOutputs.Found
	return
}

func (r *FakeRepo) Save(m app.Model) (err error) {
	r.SaveInputs.Model = m
	err = r.SaveOutputs.Err
	return
}
