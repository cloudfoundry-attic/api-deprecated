package app_test

import (
	"github.com/cloudfoundry-incubator/api/models/app"
	"github.com/cloudfoundry-incubator/api/testhelpers/test_db"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var db gorm.DB

func init() {
	db = test_db.InitTestDB()
	db.CreateTable(app.Record{})
}

var _ = Describe("App Repository", func() {
	var appRepo app.Repo

	BeforeEach(func() {
		appRepo = app.NewRepo(db)
	})

	Context("FindByGuid", func() {
		Context("when the model does not exist", func() {
			It("flags model not found", func() {
				_, found := appRepo.FindByGuid("app-1")
				Expect(found).To(Equal(false))
			})
		})

		Context("When the model exists", func() {
			It("finds the model", func() {
				model := app.NewModel()
				err := appRepo.Save(model)
				Expect(err).ToNot(HaveOccurred())

				modelFromDb, found := appRepo.FindByGuid(model.Guid())
				Expect(found).To(Equal(true))
				Expect(modelFromDb).To(Equal(model))
			})
		})
	})
})
