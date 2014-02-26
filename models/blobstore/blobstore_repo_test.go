package blobstore_test

import (
	"github.com/cloudfoundry-incubator/api/models/blobstore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"reflect"
)

var _ = Describe("BlobstoreRepo", func() {
	Context("NewRepo", func() {
		var (
			appstorePath string
			err          error
			repo         blobstore.Repo
		)

		BeforeEach(func() {
			err = nil
			appstorePath, err = ioutil.TempDir("", "blobstore_test")
			Expect(err).NotTo(HaveOccurred())

			repo = blobstore.NewRepo(blobstore.RepoArgs{
				AppPackageStore: blobstore.BlobStoreArgs{
					Filepath: appstorePath,
				},
			})
		})

		It("creates a filesytem AppPackageStore", func() {
			appStore := repo.AppPackageStore()
			isNil := reflect.ValueOf(appStore).IsNil()
			Expect(isNil).To(BeFalse())
		})
	})
})
