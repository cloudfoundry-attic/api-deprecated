package config_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry-incubator/api/config"
	"io/ioutil"
	"os"
	"path/filepath"
)

var exampleData = []byte(`
---
default_backend_url: http://www.google.com:80
port: 3000
`)

var expectedConfig = config.Config{
	DefaultBackendURL: "http://www.google.com:80",
	Port:              3000,
}

var _ = Describe("Configuration", func() {
	Context("When creating a config file from a byte slice", func() {
		It("Contains all of the set values", func() {
			c, err := config.New(exampleData)
			Expect(err).ToNot(HaveOccurred())
			Expect(c).To(Equal(expectedConfig))
		})
	})

	Context("When creating a config file from a file", func() {
		var filePath string

		BeforeEach(func() {
			dir, err := ioutil.TempDir("", "api_test")
			Expect(err).ToNot(HaveOccurred())

			filePath = filepath.Join(dir, "config.yml")

			err = ioutil.WriteFile(filePath, exampleData, os.FileMode(0600))
			Expect(err).ToNot(HaveOccurred())
		})

		AfterEach(func() {
			err := os.Remove(filePath)
			Expect(err).ToNot(HaveOccurred())
		})

		It("Contains all of the set values", func() {
			c, err := config.NewFromFile(filePath)
			Expect(err).ToNot(HaveOccurred())
			Expect(c).To(Equal(expectedConfig))
		})
	})
})
