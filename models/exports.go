package models

import (
	"github.com/cloudfoundry-incubator/api/config"
	"github.com/cloudfoundry-incubator/api/models/blobstore"
	"github.com/cloudfoundry-incubator/api/models/job"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func NewExports(db gorm.DB, c config.Config) map[interface{}]interface{} {

	return map[interface{}]interface{}{

		// JOBS
		(*job.Repo)(nil): job.NewRepo(db),

		// BLOBSTORES
		(*blobstore.Repo)(nil): blobstore.NewRepo(blobstore.RepoArgs{
			AppPackageStore: blobstore.BlobStoreArgs{
				Filepath: c.AppPackages.Filepath,
			},
		}),
	}
}
