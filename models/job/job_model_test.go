package job_test

import (
	"github.com/cloudfoundry-incubator/api/models/job"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("JobModel", func() {
	Context("Status", func() {
		It("is finished when there is no id", func() {
			j := job.NewModel()
			Expect(j.Status()).To(Equal("finished"))
		})

		It("is failed if there is an error", func() {
			r := job.NewJobRecord()
			r.LastError = "oops"

			j := job.NewModelFromRecord(r)
			Expect(j.Status()).To(Equal("failed"))
		})

		It("is queued if job has an id but no lock", func() {
			r := job.NewJobRecord()
			r.Id = int64(123)

			j := job.NewModelFromRecord(r)
			Expect(j.Status()).To(Equal("queued"))
		})

		It("is running if job has a lock and no error", func() {
			r := job.NewJobRecord()
			r.Id = int64(123)
			r.LockedAt = time.Now()

			j := job.NewModelFromRecord(r)
			Expect(j.Status()).To(Equal("running"))
		})
	})
})
