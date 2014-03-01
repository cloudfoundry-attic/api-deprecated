package database_test

import (
	"github.com/cloudfoundry-incubator/api/config"
	"github.com/cloudfoundry-incubator/api/framework/database"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Db", func() {
	var c config.DbConfig
	var db gorm.DB

	Context("sqlite", func() {
		BeforeEach(func() {
			var err error
			c = config.DbConfig{
				URI: "sqlite:///tmp/api.db",
			}
			db, err = database.NewDB(c)
			Expect(err).NotTo(HaveOccurred())
		})

		It("creates a new db conection from config", func() {
			err := db.DB().Ping()
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("mysql", func() {
		BeforeEach(func() {
			var err error
			c = config.DbConfig{
				URI: "mysql://api_user:password@/api_test?charset=utf8&parseTime=True",
			}
			db, err = database.NewDB(c)
			Expect(err).NotTo(HaveOccurred())
		})

		It("creates a new db conection from config", func() {
			err := db.DB().Ping()
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("postgres", func() {
		BeforeEach(func() {
			var err error
			c = config.DbConfig{
				URI: "postgres://postgres@localhost/api_test?sslmode=disable",
			}
			db, err = database.NewDB(c)
			Expect(err).NotTo(HaveOccurred())
		})

		It("creates a new db conection from config", func() {
			err := db.DB().Ping()
			Expect(err).NotTo(HaveOccurred())
		})
	})

})
