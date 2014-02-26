package test_db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func InitTestDB() gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	return db
}
