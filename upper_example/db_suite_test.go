package upper_example_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"upper.io/db"
	"upper.io/db/mysql"
)

var (
	session db.Database
)

func TestModel(t *testing.T) {
	var err error
	mysql.Debug = true
	session, err = db.Open("mysql", db.Settings{
		Database: "test",
		Host:     "",
		User:     "pivotal",
	})

	if err != nil {
		println("WAT", err.Error())
	}

	RegisterFailHandler(Fail)
	RunSpecs(t, "Upper Examples Suite")
}
