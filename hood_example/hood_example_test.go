package hood_example_test

import (
	"fmt"
	model "github.com/cloudfoundry-incubator/api/hood_example"
	"github.com/eaigner/hood"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HoodExample", func() {

	Context("Querying", func() {
		It("reads a single object ", func() {
			apps := []model.Apps{}
			err := hd.Where("name", "=", "simple-spring").Limit(1).Find(&apps)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(apps)).To(Equal(1))
		})

		It("finds all by relation", func() {
			apps := []model.Apps{}
			err := hd.Where("space_id", "=", 242).Find(&apps)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(apps)).ToNot(Equal(0))
		})

		It("does joins", func() {
			apps := []model.Apps{}
			err := hd.
				Join(hood.InnerJoin, "spaces", "apps.space_id", "spaces.id").
				Where("spaces.id", "=", "242").
				Find(&apps)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(apps)).ToNot(Equal(0))
		})
	})

	Context("when creating a model", func() {
		var tx *hood.Hood
		BeforeEach(func() {
			tx = hd.Begin()
		})

		AfterEach(func() {
			err := tx.Rollback()
			Expect(err).ToNot(HaveOccurred())
		})

		FIt("persists the model", func() {
			spaces := []model.Spaces{}
			err := tx.Where("id", "=", 242).Limit(1).Find(&spaces)
			Expect(err).ToNot(HaveOccurred())
			space := spaces[0]

			stacks := []model.Stacks{}
			err = tx.Where("id", "=", 1).Limit(1).Find(&stacks)
			Expect(err).ToNot(HaveOccurred())
			stack := stacks[0]

			// stack := model.Stacks{}

			app := model.Apps{
				Name:    "New guy",
				Guid:    model.GenerateUUID(),
				SpaceId: int64(space.Id),
				StackId: int64(stack.Id),
			}

			ids, err := tx.SaveAll(&[]model.Apps{app})
			Expect(err).ToNot(HaveOccurred())
			fmt.Printf("%#v", ids)

			// tx = tx.Save(app)

			// Expect(tx.Error).To(BeNil())
			// Expect(app.Id).ToNot(Equal(0))

			// newApp := new(model.App)
			// err := tx.First(newApp, app.Id).
			// 	Update(map[string]interface{}{"name": "Better name", "age": 18}).Error
			// Expect(err).To(BeNil())

			// Expect(newApp.Name).To(Equal("Better name"))
			// Expect(newApp.SpaceId).To(Equal(space.Id))
			// Expect(newApp.StackId).To(Equal(stack.Id))
		})
	})
})
