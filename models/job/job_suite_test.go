package job_test

import (
	"github.com/cloudfoundry-incubator/api/models/job"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var db gorm.DB

func TestJob(t *testing.T) {
	var err error
	RegisterFailHandler(Fail)
	db, err = gorm.Open("sqlite3", ":memory:")
	db.LogMode(true)
	Expect(err).ToNot(BeNil())
	Expect(db).ToNot(BeNil())
	db.CreateTable(job.JobRecord{})
	RunSpecs(t, "Job Suite")
}
