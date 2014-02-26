package middle_test

import (
	. "github.com/cloudfoundry-incubator/api/testhelpers/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http/httptest"

	"github.com/cloudfoundry-incubator/api/framework/json"
	"github.com/cloudfoundry-incubator/api/framework/middle"
	"github.com/codegangsta/martini"
	"io/ioutil"
	"net/http"
)

var _ = Describe("Response Handler", func() {
	var m *martini.ClassicMartini
	var s *httptest.Server

	BeforeEach(func() {
		m = martini.Classic()
		m.Use(middle.ResponseHandler)
		s = httptest.NewServer(m)
	})

	AfterEach(func() {
		s.Close()
	})

	It("injects the response object into the request", func() {
		expectedJSON := `
			{"foo":"bar"}
		`
		m.Get("**", func(res middle.Response) {
			res.RenderJson(201, json.Map{
				"foo": "bar",
			})
		})

		resp, err := http.Get(s.URL)
		Expect(err).NotTo(HaveOccurred())

		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(body)).To(MatchJson(expectedJSON))
		Expect(resp.StatusCode).To(Equal(201))
	})
})
