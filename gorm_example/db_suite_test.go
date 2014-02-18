package model_test

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"os"
	"testing"
)

var db gorm.DB

func TestModel(t *testing.T) {

	var err error
	db, err = gorm.Open("mysql", "pivotal:@/test?charset=utf8&parseTime=True")
	if err != nil {
		println("WAT", err.Error())
	}

	logger := log.New(os.Stdout, "\r\n", 0)
	db.SetLogger(gorm.Logger{logger})
	db.LogMode(true)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Model Suite")
}
