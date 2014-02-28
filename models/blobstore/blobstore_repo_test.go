package blobstore_test

import (
	"fmt"
	"github.com/cloudfoundry-incubator/api/models/blobstore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"reflect"
)

var _ = Describe("BlobstoreRepo", func() {
	Context("NewRepo", func() {
		var (
			err              error
			repo             blobstore.Repo
			appBlobstoreArgs blobstore.BlobStoreArgs
		)

		BeforeEach(func() {
			err = nil
		})

		Context("when the provider is local", func() {
			var (
				appstorePath string
			)
			BeforeEach(func() {
				appstorePath, err = ioutil.TempDir("", "blobstore_test")
				Expect(err).NotTo(HaveOccurred())

				appBlobstoreArgs = blobstore.BlobStoreArgs{
					Provider: "local",
					Filepath: appstorePath,
				}
				repoArgs := blobstore.RepoArgs{AppPackageStore: appBlobstoreArgs}
				repo = blobstore.NewRepo(repoArgs)
			})

			It("creates a filesytem AppPackageStore", func() {
				appStore := repo.AppPackageStore()
				isNil := reflect.ValueOf(appStore).IsNil()
				Expect(isNil).To(BeFalse())
				formattedValue := fmt.Sprintf("%s", reflect.ValueOf(appStore))
				Expect(formattedValue).To(MatchRegexp("blobstore.fileStore"))
			})
		})

		Context("When provider is s3", func() {
			BeforeEach(func() {
				appBlobstoreArgs = blobstore.BlobStoreArgs{
					Provider:        "s3",
					AccessKeyId:     "key_id",
					AccessKeySecret: "key_secret",
					Host:            "s3.amazon.com",
					BucketName:      "app_package",
				}

				repoArgs := blobstore.RepoArgs{AppPackageStore: appBlobstoreArgs}
				repo = blobstore.NewRepo(repoArgs)
			})

			It("creates a s3 AppPackageStore", func() {
				appStore := repo.AppPackageStore()
				isNil := reflect.ValueOf(appStore).IsNil()
				Expect(isNil).To(BeFalse())
				formattedValue := fmt.Sprintf("%s", reflect.ValueOf(appStore))
				Expect(formattedValue).To(MatchRegexp("blobstore.s3store"))
			})
		})
	})
})
