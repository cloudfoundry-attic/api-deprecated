package middle_test

import (
	"github.com/cloudfoundry-incubator/api/framework/middle"
	"github.com/codegangsta/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"net/url"
)

var _ = Describe("Request Handler", func() {
	var m *martini.ClassicMartini
	var s *httptest.Server

	BeforeEach(func() {
		m = martini.Classic()
		m.Use(middle.RequestHandler)
		s = httptest.NewServer(m)
	})

	AfterEach(func() {
		s.Close()
	})

	It("injects the request object into the request", func() {
		var value string

		m.Post("**", func(req middle.Request) {
			value = req.Param("foo")
		})

		_, err := http.PostForm(s.URL, url.Values{"foo": {"bar"}})
		Expect(err).NotTo(HaveOccurred())
		Expect(value).To(Equal("bar"))
	})
})
