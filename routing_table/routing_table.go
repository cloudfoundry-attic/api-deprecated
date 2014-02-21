package routing_table

import (
	"github.com/cloudfoundry-incubator/api/controllers/job_controller"
	"github.com/cloudfoundry-incubator/api/framework/router"
)

func New() []router.Route {
	return []router.Route{
		{
			Method: "get", Path: "/v2/jobs/:guid", Action: job_controller.Get,
		},
	}
}
