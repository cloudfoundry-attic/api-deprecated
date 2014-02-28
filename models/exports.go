package models

import (
	"github.com/cloudfoundry-incubator/api/config"
	"github.com/cloudfoundry-incubator/api/models/app"
	"github.com/cloudfoundry-incubator/api/models/blobstore"
	"github.com/cloudfoundry-incubator/api/models/job"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func NewExports(db gorm.DB, c config.Config) map[interface{}]interface{} {
	return map[interface{}]interface{}{

		// JOBS
		(*job.Repo)(nil): job.NewRepo(db),

		// APPS
		(*app.Repo)(nil): app.NewRepo(db),

		// BLOBSTORES
		(*blobstore.Repo)(nil): blobstore.NewRepo(blobstore.RepoArgs{
			AppPackageStore: blobstore.BlobStoreArgs{
				Provider:        blobstore.Provider(c.AppPackages.Provider),
				Filepath:        c.AppPackages.Filepath,
				AccessKeyId:     c.AppPackages.AccessKeyId,
				AccessKeySecret: c.AppPackages.AccessKeySecret,
				Host:            c.AppPackages.Host,
				BucketName:      c.AppPackages.BucketName,
			},
		}),
	}
}
