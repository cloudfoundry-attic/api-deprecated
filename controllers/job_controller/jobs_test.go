package job_controller_test

import (
	"github.com/cloudfoundry-incubator/api/controllers/job_controller"
	middle "github.com/cloudfoundry-incubator/api/middleware/http"
	. "github.com/cloudfoundry-incubator/api/testhelpers/matchers"
	"github.com/cloudfoundry-incubator/api/testhelpers/models/fake_job"
	"github.com/codegangsta/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"time"
)

func StringToTime(s string) time.Time {
	t, err := time.Parse("2006-01-02T15:04:05+00:00", s)
	if err != nil {
		panic("time parsing error: " + err.Error())
	}
	println("TIME", t.Format("2006-01-02T15:04:05+00:00"))
	return t
}

var _ = Describe("Jobs", func() {
	var res *middle.TestResponse
	var req *middle.TestRequest

	Context("Get", func() {
		Context("when the job exists", func() {
			BeforeEach(func() {
				res = middle.NewTestResponse()
				req = middle.NewTestRequest()
				params := martini.Params{}

				jobModel := &fake_job.Model{}
				jobModel.Outputs.Guid = "job-guid"
				jobModel.Outputs.CreatedAt = StringToTime("2014-02-10T05:38:46+00:00")
				jobModel.Outputs.Url = "/v2/jobs/job-guid"
				jobModel.Outputs.Status = "queued"

				jobRepo := &fake_job.Repo{}
				jobRepo.FindByGuidOutput.Model = jobModel

				job_controller.Get(res, req, params, jobRepo)
			})

			It("returns 200 OK", func() {
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})

			It("returns job as JSON", func() {
				expectedJSON := `
					{
					  "metadata": {
					    "guid": "job-guid",
					    "created_at": "2014-02-10T05:38:46+00:00",
					    "url": "/v2/jobs/job-guid"
					  },
					  "entity": {
					    "guid": "job-guid",
					    "status": "queued"
					  }
					}
				`
				Expect(res.Body).To(MatchJson(expectedJSON))
			})
		})
	})
})
