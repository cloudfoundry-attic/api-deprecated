package json_test

import (
	"github.com/cloudfoundry-incubator/api/framework/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Json", func() {
	var (
		timestamp      time.Time
		expectedFormat = "2006-01-02T15:04:05+00:00"
	)

	BeforeEach(func() {
		var err error
		timestamp, err = time.Parse("2006-01-02T15:04:05+00:00", "2006-01-02T15:04:05+00:00")
		Expect(err).NotTo(HaveOccurred())

	})
	It("converts time.Time to Time with correct format", func() {
		jsonObject := json.Map{
			"time": timestamp,
		}

		jsonData, err := jsonObject.MarshalJSON()
		Expect(err).NotTo(HaveOccurred())
		Expect(string(jsonData)).To(ContainSubstring(expectedFormat))
	})

	It("deals with nested maps", func() {
		jsonObject := json.Map{
			"obj": json.Map{
				"timestamp": timestamp,
			},
		}
		jsonData, err := jsonObject.MarshalJSON()
		Expect(err).NotTo(HaveOccurred())
		Expect(string(jsonData)).To(ContainSubstring(expectedFormat))
	})

	It("deals with nested slices", func() {
		jsonObject := json.Map{
			"obj": json.Array{
				timestamp,
			},
		}
		jsonData, err := jsonObject.MarshalJSON()
		Expect(err).NotTo(HaveOccurred())
		Expect(string(jsonData)).To(ContainSubstring(expectedFormat))
	})
})
