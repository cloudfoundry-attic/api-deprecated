package model_test

import (
	model "github.com/cloudfoundry-incubator/api/gorm_example"

	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gorm Spike", func() {

	var tx *gorm.DB

	BeforeEach(func() {
		println("Begin")
		tx = db.Begin()
	})

	AfterEach(func() {
		println("Rolling back")

		tx.Rollback()
	})

	Context("When querying", func() {
		It("reads a single object", func() {
			app := new(model.App)
			tx = tx.First(app)
			Expect(tx.Error).To(BeNil())
			Expect(app.Name).ToNot(BeEmpty())
		})

		It("find all by relation", func() {
			space := new(model.Space)
			tx = tx.First(space)

			apps := new([]model.App)
			tx1 := tx.Model(space)
			tx1.Related(apps) //WAAAAAAAAT
			Expect(tx1.Error).To(BeNil())
			Expect(len(*apps)).ToNot(Equal(0))
		})
	})

	Context("When creating a model", func() {
		var space *model.Space
		var stack *model.Stack

		BeforeEach(func() {
			space = new(model.Space)
			stack = new(model.Stack)

			tx = tx.First(space).First(stack)
			Expect(tx.Error).To(BeNil())
		})

		FIt("persists the model", func() {
			app := &model.App{
				Name:    "New guy",
				Guid:    model.GenerateUUID(),
				SpaceId: space.Id,
				StackId: stack.Id,
			}

			tx = tx.Save(app)

			Expect(tx.Error).To(BeNil())
			Expect(app.Id).ToNot(Equal(0))

			newApp := new(model.App)
			err := tx.First(newApp, app.Id).
				Update(map[string]interface{}{"name": "Better name", "age": 18}).Error
			Expect(err).To(BeNil())

			Expect(newApp.Name).To(Equal("Better name"))
			Expect(newApp.SpaceId).To(Equal(space.Id))
			Expect(newApp.StackId).To(Equal(stack.Id))
		})
	})
})
