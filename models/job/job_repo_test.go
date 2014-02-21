package job_test

import (
	"github.com/cloudfoundry-incubator/api/models/job"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("JobRepo", func() {
	var jobRepo job.JobRepo
	BeforeEach(func() {
		jobRepo = job.NewRepo(db)

	})

	Context("FindByGuid", func() {
		Context("when the model does not exist", func() {
			It("flags model not found", func() {
				_, found := jobRepo.FindByGuid("job-1")
				Expect(found).To(Equal(false))
			})
		})

		Context("When the model exists", func() {
			It("finds the model", func() {
				model := job.NewModel()
				err := jobRepo.Save(model)
				Expect(err).ToNot(HaveOccurred())

				modelFromDb, found := jobRepo.FindByGuid(model.Guid())
				Expect(found).To(Equal(true))
				Expect(modelFromDb).To(Equal(model))
			})
		})
	})

	Context("Save", func() {
		It("sets the timezone to UTC on last updated at", func() {
			r := job.NewJobRecord()
			r.LockedAt = time.Now()

			model := job.NewModelFromRecord(r)
			err := jobRepo.Save(model)
			Expect(err).ToNot(HaveOccurred())

			modelFromDb, found := jobRepo.FindByGuid(model.Guid())
			Expect(found).To(Equal(true))
			Expect(modelFromDb).To(Equal(model))
		})
	})
})
