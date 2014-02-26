package app_test

import (
	"github.com/cloudfoundry-incubator/api/models/app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App Model", func() {
	var (
		a app.Model
	)

	Context("SetPackageHash", func() {
		BeforeEach(func() {
			a = app.NewModel()
			a.SetPackageHash("abc")
		})

		It("sets the package hash", func() {
			Expect(a.PackageHash()).To(Equal("abc"))
		})

		It("sets the package state to pending", func() {
			Expect(a.PackageState()).To(Equal("PENDING"))
		})
	})
})
