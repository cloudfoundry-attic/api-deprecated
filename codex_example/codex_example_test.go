package codex_example_test

import (
	"fmt"
	"github.com/chuckpreslar/codex"
	"github.com/chuckpreslar/codex/managers"
	model "github.com/cloudfoundry-incubator/api/codex_example"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Arel in Go", func() {
	var appTable managers.Accessor

	BeforeEach(func() {
		appTable = codex.Table("apps")
	})

	Context("Query objects", func() {
		It("reads a single object", func() {
			query := appTable.Project(model.AppKeys...).Where(appTable("name").Eq("simple-spring")).Limit(1)
			query.SetAdapter("mysql") // can we move this to a higher level object?

			sql, err := query.ToSql()
			Expect(err).To(BeNil())
			fmt.Printf(sql)

			vals := []interface{}{new(int64), new(string), new(string), new(int64), new(int64)}
			row := db.QueryRow(sql)
			err = row.Scan(vals...)
			Expect(err).To(BeNil())

			app := model.ZipMap(model.AppKeys, vals)

			fmt.Printf("\nAPP: %#v\n", app)

			Expect(app["id"]).ToNot(Equal(int64(0)))
			Expect(app["name"]).ToNot(Equal("simple-spring"))
		})

		It("find by name", func() {
		})
	})

	Context("Update objects", func() {
		It("updates app name", func() {
			query := appTable.Project("id").Where(appTable("name").Eq("nyet-test")).Limit(1)
		})
	})
})
