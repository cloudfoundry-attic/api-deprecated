package routing_table

import (
	"github.com/cloudfoundry-incubator/api/controllers/job_controller"
	"github.com/cloudfoundry-incubator/api/router"
)

func New() []router.Route {
	return []router.Route{
		{
			Method: "get", Path: "/v2/jobs", Action: job_controller.Get,
		},
	}
}
