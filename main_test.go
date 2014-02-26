package main_test

import (
	testnet "github.com/cloudfoundry-incubator/api/testhelpers/net"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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

	Context("PUT /v2/apps/:app_guid/bits", func() {
		var (
			res *http.Response
		)

		BeforeEach(func() {
			cwd, err := os.Getwd()
			Expect(err).NotTo(HaveOccurred())

			fixturePath := filepath.Join(cwd, "test_fixtures/dora.zip")
			formValues := url.Values{
				"application_path": {fixturePath},
			}
			body := strings.NewReader(formValues.Encode())

			req, err := http.NewRequest("PUT", proxyURL+"/v2/apps/app-guid-1/bits", body)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			Expect(err).NotTo(HaveOccurred())

			res, err = http.DefaultClient.Do(req)
			Expect(err).NotTo(HaveOccurred())
		})

		It("uploads the file", func() {
			Expect(res.StatusCode).To(Equal(http.StatusOK))

			expectedFile := filepath.Join(conf.AppPackages.Filepath, "app-guid-1")
			_, err := os.Stat(expectedFile)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
