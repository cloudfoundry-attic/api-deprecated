package job_test

import (
	"github.com/cloudfoundry-incubator/api/models/job"
	"github.com/cloudfoundry-incubator/api/testhelpers/test_db"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var db gorm.DB

func init() {
	db = test_db.InitTestDB()
	db.CreateTable(job.Record{})
}

var _ = Describe("Repo", func() {
	var jobRepo job.Repo
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
			r := job.NewRecord()
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
