package intercept_proxy_test

import (
	"fmt"
	"github.com/cloudfoundry-incubator/api/intercept_proxy"
	testnet "github.com/cloudfoundry-incubator/api/testhelpers/net"
	. "github.com/onsi/ginkgo"
	gconfig "github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Intercept Proxy", func() {
	var (
		p              intercept_proxy.Proxy
		defaultBackend *httptest.Server
		defaultHandler *testnet.TestHandler
		port           int
	)
	BeforeEach(func() {
		port = 3000 + gconfig.GinkgoConfig.ParallelNode

		defaultBackend, defaultHandler = testnet.NewServer(GinkgoT(), []testnet.TestRequest{})
		p = intercept_proxy.New(intercept_proxy.Args{
			DefaultBackendURL: defaultBackend.URL,
		})

		proxyServer := &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: p,
		}

		go proxyServer.ListenAndServe()
	})

	Context("when the route does not exist", func() {
		It("should proxy to the default backend", func() {
			defaultHandler.Requests = []testnet.TestRequest{
				{
					Method:   "GET",
					Path:     "/foo",
					Response: testnet.TestResponse{Status: http.StatusOK, Body: "hello world"}}}

			resp, err := http.DefaultClient.Get(fmt.Sprintf("http://localhost:%d/foo", port))
			Expect(err).ToNot(HaveOccurred())

			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(body)).To(Equal("hello world\n"))
		})
	})

	Context("when the route does not exist ad", func() {
		It("should proxy to the default backend asd", func() {
			defaultHandler.Requests = []testnet.TestRequest{
				{
					Method:   "get",
					Path:     "/foo",
					Response: testnet.TestResponse{Status: http.StatusOK, Body: "hello world"}}}

			resp, err := http.DefaultClient.Get(fmt.Sprintf("http://localhost:%d/foo", port))
			Expect(err).ToNot(HaveOccurred())

			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(body)).To(Equal("hello world\n"))
		})
	})
})
