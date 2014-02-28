package main_test

import (
	testnet "github.com/cloudfoundry-incubator/api/testhelpers/net"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
)

var _ = Describe("API INTEGRATION RUNNER", func() {
	Context("when the route does not exist", func() {

		BeforeEach(func() {
			defaultHandler.Requests = []testnet.TestRequest{{
				Method:   "GET",
				Path:     "/foo",
				Response: testnet.TestResponse{Status: http.StatusOK, Body: "hello world"},
			}}
		})

		It("should proxy to the default backend", func() {
			resp, err := http.DefaultClient.Get(proxyURL + "/foo")

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(defaultHandler.AllRequestsCalled()).To(BeTrue())

			body, err := ioutil.ReadAll(resp.Body)
			Expect(string(body)).To(Equal("hello world\n"))
		})
	})

	Context("when the route does exist", func() {
		It("handles the request", func() {
			res, err := http.DefaultClient.Get(proxyURL + "/v2/jobs/abcdef")
			Expect(err).ToNot(HaveOccurred())
			Expect(res.StatusCode).To(Equal(http.StatusOK))
		})
	})
})
