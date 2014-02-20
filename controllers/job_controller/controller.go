package job_controller

import (
	middle "github.com/cloudfoundry-incubator/api/middleware/http"
	"github.com/cloudfoundry-incubator/api/models/job"
	"github.com/codegangsta/martini"
	"net/http"
)

const TIMESTAMP_FORMAT = "2006-01-02T15:04:05+00:00"

func Get(res middle.Response, req middle.Request, params martini.Params, jobRepo job.Repo) {
	res.SetStatusCode(http.StatusOK)
	j := jobRepo.FindByGuid(params["guid"])
	res.WriteJson(NewJobResource(j))
}

func NewJobResource(j job.Model) JobResource {
	return JobResource{
		Metadata: JobMetadata{
			Guid:      j.Guid(),
			CreatedAt: j.CreatedAt().Format(TIMESTAMP_FORMAT),
			Url:       j.Url(),
		},
		Entity: JobEntity{
			Guid:   j.Guid(),
			Status: j.Status(),
		},
	}
}

type JobResource struct {
	Metadata JobMetadata `json:"metadata"`
	Entity   JobEntity   `json:"entity"`
}

type JobMetadata struct {
	Guid      string `json:"guid"`
	CreatedAt string `json:"created_at"`
	Url       string `json:"url"`
}

type JobEntity struct {
	Guid   string `json:"guid"`
	Status string `json:"status"`
}
