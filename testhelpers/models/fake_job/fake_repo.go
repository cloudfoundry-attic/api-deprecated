package fake_job

import "github.com/cloudfoundry-incubator/api/models/job"

type Repo struct {
	FindByGuidInput struct {
		Guid string
	}
	FindByGuidOutput struct {
		Model *Model
	}
}

func (repo *Repo) FindByGuid(guid string) job.Model {
	repo.FindByGuidInput.Guid = guid
	return repo.FindByGuidOutput.Model
}
