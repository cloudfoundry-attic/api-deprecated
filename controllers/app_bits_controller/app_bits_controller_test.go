package app_bits_controller_test

import (
	"github.com/cloudfoundry-incubator/api/controllers/app_bits_controller"
	"github.com/cloudfoundry-incubator/api/digest"
	"github.com/cloudfoundry-incubator/api/models/app"
	"github.com/cloudfoundry-incubator/api/testhelpers/file"
	"github.com/cloudfoundry-incubator/api/testhelpers/middleware"
	"github.com/cloudfoundry-incubator/api/testhelpers/models/fake_app"
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
			fixturePath       = filepath.Join(file.Cwd(), "../../test_fixtures/dora.zip")
			req               *middleware.TestRequest
			res               *middleware.TestResponse
			fakeBlobstoreRepo *fake_blobstore.Repo
			fakeBlobstore     *fake_blobstore.Blobstore
			expectedBytes     []byte
			tmpPath           string
			appModel          app.Model
			fakeAppRepo       *fake_app.FakeRepo
		)

		BeforeEach(func() {
			var err error

			tmpPath = file.TmpFile()
			file.Copy(fixturePath, tmpPath)

			expectedBytes, err = ioutil.ReadFile(tmpPath)
			Expect(err).NotTo(HaveOccurred())

			req = middleware.NewTestRequest()
			req.Params = map[string]string{
				"application_path": tmpPath,
			}

			res = middleware.NewTestResponse()

			params := martini.Params{
				"app_guid": "app-guid-1",
			}

			fakeBlobstore = fake_blobstore.NewFakeBlobstore()
			fakeBlobstoreRepo = fake_blobstore.NewFakeBlobstoreRepo()
			fakeBlobstoreRepo.AppPackageBlobstore = fakeBlobstore

			appModel = app.NewModel()
			fakeAppRepo = fake_app.NewFakeRepo()
			fakeAppRepo.FindByGuidOutputs.Model = appModel
			fakeAppRepo.FindByGuidOutputs.Found = true

			app_bits_controller.Put(req, res, params, fakeBlobstoreRepo, fakeAppRepo)
		})

		It("returns 201 CREATED", func() {
			Expect(res.StatusCode).To(Equal(http.StatusCreated))
		})

		It("returns empty json", func() {
			Expect(res.Body).To(Equal("{}"))
		})

		It("uploads the file to blobstore", func() {
			Expect(fakeBlobstore.UploadInputs.Key).To(Equal("app-guid-1"))
			Expect(fakeBlobstore.UploadInputs.File).To(Equal(expectedBytes))
		})

		It("removes the uploaded files", func() {
			_, err := os.Stat(tmpPath)
			Expect(err).To(HaveOccurred())
		})

		It("fetches the app by guid", func() {
			Expect(fakeAppRepo.FindByGuidInputs.Guid).To(Equal("app-guid-1"))
		})

		It("updates the app package hash", func() {
			f, err := os.Open(fixturePath)
			defer f.Close()
			Expect(err).NotTo(HaveOccurred())

			hexdigest, err := digest.Hex(f)
			Expect(err).NotTo(HaveOccurred())

			Expect(appModel.PackageHash()).To(Equal(hexdigest))
		})

		It("saves the app", func() {
			Expect(fakeAppRepo.SaveInputs.Model).To(Equal(appModel))
		})
	})
})
