package router_test

import (
	"fmt"
	"github.com/cloudfoundry-incubator/api/router"
	testnet "github.com/cloudfoundry-incubator/api/testhelpers/net"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tjarratt/mr_t"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe("Router", func() {
	var (
		r              router.Router
		defaultBackend *httptest.Server
		defaultHandler *testnet.TestHandler
		routes         []router.Route
		server         *httptest.Server
	)

	JustBeforeEach(func() {

		defaultBackend, defaultHandler = testnet.NewServer(mr_t.T(), []testnet.TestRequest{})

		r = router.New(router.Args{
			DefaultBackendURL: defaultBackend.URL,
			Routes:            routes,
		})

		server = httptest.NewServer(r)
	})

	AfterEach(func() {
		server.Close()
	})

	Context("when the route does not exist", func() {
		It("proxies to the default backend", func() {
			defaultHandler.Requests = []testnet.TestRequest{
				{
					Method:   "GET",
					Path:     "/foo/bar?id=2",
					Response: testnet.TestResponse{Status: http.StatusOK, Body: "hello world"}}}

			resp, err := http.DefaultClient.Get(fmt.Sprintf("%s/foo/bar?id=2", server.URL))
			Expect(err).ToNot(HaveOccurred())
			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(body)).To(Equal("hello world\n"))
		})
	})

	Context("when the route does exist", func() {
		var (
			getReq *http.Request
			get    = func(response http.ResponseWriter, request *http.Request) {
				getReq = request
			}

			postReq *http.Request
			post    = func(response http.ResponseWriter, request *http.Request) {
				postReq = request
			}
		)

		BeforeEach(func() {
			getReq = nil
			postReq = nil

			routes = []router.Route{
				{
					Method: "GET",
					Path:   "/jobs",
					Action: get,
				},
				{
					Method: "POST",
					Path:   "/jobs",
					Action: post,
				},
			}
		})

		It("routes the request to the action", func() {
			_, err := http.DefaultClient.Get(fmt.Sprintf("%s/jobs", server.URL))
			Expect(err).ToNot(HaveOccurred())
			Expect(getReq).ToNot(BeNil())
		})

		It("matches the verb", func() {
			body := strings.NewReader("foo")
			_, err := http.DefaultClient.Post(fmt.Sprintf("%s/jobs", server.URL), "text", body)
			Expect(err).ToNot(HaveOccurred())
			Expect(postReq).ToNot(BeNil())
			Expect(getReq).To(BeNil())
		})
	})
})
