package routing_table

import (
	"github.com/cloudfoundry-incubator/api/controllers/app_bits_controller"
	"github.com/cloudfoundry-incubator/api/controllers/job_controller"
	"github.com/cloudfoundry-incubator/api/framework/router"
)

func New() []router.Route {
	return []router.Route{
		{Method: "get", Path: "/v2/jobs/:job_guid", Action: job_controller.Get},
		{Method: "put", Path: "/v2/apps/:app_guid/bits", Action: app_bits_controller.Put},
	}
}
