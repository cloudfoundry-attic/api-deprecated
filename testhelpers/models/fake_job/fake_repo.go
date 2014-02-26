package fake_job

import "github.com/cloudfoundry-incubator/api/models/job"

type Repo struct {
	FindByGuidInput struct {
		Guid string
	}
	FindByGuidOutput struct {
		Model *Model
		Found bool
	}
}

func (repo *Repo) FindByGuid(guid string) (job.Model, bool) {
	repo.FindByGuidInput.Guid = guid
	return repo.FindByGuidOutput.Model, repo.FindByGuidOutput.Found
}

func (repo *Repo) Save(job.Model) error {
	return nil
}
