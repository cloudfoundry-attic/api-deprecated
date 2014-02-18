package codex_example_test

import (
	"database/sql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

var db *sql.DB

func TestCodexExample(t *testing.T) {
	var err error
	db, err = sql.Open("mysql", "pivotal:@/test?charset=utf8&parseTime=True")
	if err != nil {
		println("WAT", err.Error())
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "Arel Suite")
}
