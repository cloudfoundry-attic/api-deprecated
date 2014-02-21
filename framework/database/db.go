package database

import (
	"errors"
	"github.com/cloudfoundry-incubator/api/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"net/url"
	"strings"
)

func NewDB(c config.DbConfig) (db gorm.DB, err error) {
	dbURL, err := url.Parse(c.URI)
	if err != nil {
		return
	}

	var dbType, dbParams string
	switch dbURL.Scheme {
	case "sqlite":
		dbType = "sqlite3"
		dbParams = dbURL.Path
	case "postgres":
		dbType = "postgres"
		dbParams = dbURL.String()
	case "mysql":
		dbType = "mysql"
		dbParams = strings.TrimLeft(dbURL.RequestURI(), "/")
	default:
		err = errors.New("Unsupported db type: " + dbURL.Scheme)
		return
	}

	db, err = gorm.Open(dbType, dbParams)
	if err != nil {
		return
	}
	db.LogMode(true)

	err = db.DB().Ping()
	if err != nil {
		return
	}

	return
}
