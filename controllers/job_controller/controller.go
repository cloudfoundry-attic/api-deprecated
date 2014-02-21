package job_controller

import (
	"github.com/cloudfoundry-incubator/api/framework/json"
	middle "github.com/cloudfoundry-incubator/api/framework/middleware/http"
	"github.com/cloudfoundry-incubator/api/models/job"
	"github.com/codegangsta/martini"
	"net/http"
)

func Get(res middle.Response, params martini.Params, jobRepo job.JobRepo) {
	j, found := jobRepo.FindByGuid(params["guid"])
	if !found {
		j = job.NewModel()
	}
	res.RenderJson(http.StatusOK, NewJobResource(j))
}

func NewJobResource(j job.JobModel) json.Resource {
	return json.Resource{
		Metadata: json.Map{
			"guid":       j.Guid(),
			"created_at": j.CreatedAt(),
			"url":        j.Url(),
		},
		Entity: json.Map{
			"guid":   j.Guid(),
			"status": j.Status(),
		},
	}
}
