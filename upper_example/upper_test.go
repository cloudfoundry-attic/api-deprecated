package upper_example_test

import (
	"fmt"
	model "github.com/cloudfoundry-incubator/api/upper_example"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"upper.io/db"
)

var _ = Describe("Gorm Spike", func() {
	var appCollection, spaceCollection db.Collection

	BeforeEach(func() {
		var err error
		appCollection, err = session.Collection("apps")
		Expect(err).ToNot(HaveOccurred())
		spaceCollection, err = session.Collection("spaces")
		Expect(err).ToNot(HaveOccurred())
	})

	Context("When querying", func() {
		It("reads a single object", func() {
			res := appCollection.Find().Limit(1)
			app := new(model.App)
			err := res.One(app)

			Expect(app.Name).ToNot(BeEmpty())
			Expect(err).ToNot(HaveOccurred())
		})

		It("find all by relation", func() {
			space := new(model.Space)
			spaceCollection.Find().Limit(1).One(space)

			res := appCollection.Find(db.Cond{"space_id": space.Id})
			apps := new([]model.App)
			err := res.All(apps)

			fmt.Printf("%#v", apps)
			Expect(len(*apps)).ToNot(Equal(0))
			Expect(err).ToNot(HaveOccurred())
		})

		It("can do joins", func() {

		})
	})

	// Context("When creating a model", func() {
	// 	var space *model.Space
	// 	var stack *model.Stack

	// BeforeEach(func() {
	// 	space = new(model.Space)
	// 	stack = new(model.Stack)

	// 	tx = tx.First(space).First(stack)
	// 	Expect(tx.Error).To(BeNil())
	// })

	// 	FIt("persists the model", func() {
	// 		app := &model.App{
	// 			Name:    "New guy",
	// 			Guid:    model.GenerateUUID(),
	// 			SpaceId: space.Id,
	// 			StackId: stack.Id,
	// 		}

	// 		tx = tx.Save(app)

	// 		Expect(tx.Error).To(BeNil())
	// 		Expect(app.Id).ToNot(Equal(0))

	// 		newApp := new(model.App)
	// 		err := tx.First(newApp, app.Id).
	// 			Update(map[string]interface{}{"name": "Better name", "age": 18}).Error
	// 		Expect(err).To(BeNil())

	// 		Expect(newApp.Name).To(Equal("Better name"))
	// 		Expect(newApp.SpaceId).To(Equal(space.Id))
	// 		Expect(newApp.StackId).To(Equal(stack.Id))
	// 	})
	// })
})
