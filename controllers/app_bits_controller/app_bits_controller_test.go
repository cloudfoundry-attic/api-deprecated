package app_bits_controller_test

import (
	"github.com/cloudfoundry-incubator/api/controllers/app_bits_controller"
	"github.com/cloudfoundry-incubator/api/testhelpers/middleware"
	"github.com/cloudfoundry-incubator/api/testhelpers/models/fake_blobstore"
	"github.com/codegangsta/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var _ = Describe("AppBitsController", func() {
	Describe("PUT app bits", func() {
		var (
			req               *middleware.TestRequest
			res               *middleware.TestResponse
			fakeBlobstoreRepo *fake_blobstore.Repo
			fakeBlobstore     *fake_blobstore.Blobstore
			expectedBytes     []byte
		)

		BeforeEach(func() {
			cwd, err := os.Getwd()
			Expect(err).NotTo(HaveOccurred())
			fixturePath := filepath.Join(cwd, "../../test_fixtures/dora.zip")

			expectedBytes, err = ioutil.ReadFile(fixturePath)
			Expect(err).NotTo(HaveOccurred())

			req = middleware.NewTestRequest()
			req.Params = map[string]string{
				"application_path": fixturePath,
			}

			res = middleware.NewTestResponse()

			params := martini.Params{
				"app_guid": "app-guid-1",
			}

			fakeBlobstore = fake_blobstore.NewFakeBlobstore()

			fakeBlobstoreRepo = fake_blobstore.NewFakeBlobstoreRepo()
			fakeBlobstoreRepo.AppPackageBlobstore = fakeBlobstore

			app_bits_controller.Put(req, res, params, fakeBlobstoreRepo)
		})

		It("returns 200 OK", func() {
			Expect(res.StatusCode).To(Equal(http.StatusOK))
		})

		It("uploads the file to blobstore", func() {
			Expect(fakeBlobstore.UploadInputs.Key).To(Equal("app-guid-1"))
			Expect(fakeBlobstore.UploadInputs.File).To(Equal(expectedBytes))
		})

		XIt("updates the app package hash")
		XIt("removes the uploaded files")
	})
})
