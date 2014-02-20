package http_test

import (
	"fmt"
	middle "github.com/cloudfoundry-incubator/api/middleware/http"
	"github.com/codegangsta/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Request", func() {
	var m *martini.ClassicMartini

	BeforeEach(func() {
		m = martini.Classic()
		m.Use(middle.RequestHandler)
	})

	It("creates a new request", func() {
		expectedParams := map[string]string{
			"action": "foo",
			"id":     "123",
		}
		var actualReq middle.Request
		m.Get("/:action/:id", func(req http.Request) {
			fmt.Printf("%#v", parma)
		})

		server := httptest.NewServer(m)
		_, err := http.Get(fmt.Sprintf("%s/foo/123", server.URL))
		Expect(err).NotTo(HaveOccurred())
		Expect(actualReq.Params()).To(Equal(expectedParams))
		server.Close()
	})
})
