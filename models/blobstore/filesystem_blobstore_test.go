package blobstore_test

import (
	"github.com/cloudfoundry-incubator/api/models/blobstore"
	"github.com/cloudfoundry-incubator/api/testhelpers/file"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
	"path/filepath"
)

var _ = Describe("File System Blobstore", func() {
	var (
		blobstorePath string
		err           error
		fileStore     blobstore.BlobStore
	)

	BeforeEach(func() {
		err = nil
		blobstorePath, err = ioutil.TempDir("", "blobstore_test")
		Expect(err).NotTo(HaveOccurred())
		fileStore = blobstore.NewFileSystemBlobStore(blobstorePath)
	})

	Context("Upload", func() {
		var (
			fixturePath1 = filepath.Join(file.Cwd(), "../../test_fixtures/dora.zip")
			fixturePath2 = filepath.Join(file.Cwd(), "../../test_fixtures/app.zip")
			fixture1     *os.File
			fixture2     *os.File
		)

		BeforeEach(func() {
			fixture1, err = os.Open(fixturePath1)
			Expect(err).NotTo(HaveOccurred())
			fixture2, err = os.Open(fixturePath2)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			fixture1.Close()
			fixture2.Close()
		})

		It("writes the file to the correct location", func() {
			err = fileStore.Upload("uploaded_file", fixture1)
			Expect(err).NotTo(HaveOccurred())

			uploadSize := fileSize(filepath.Join(blobstorePath, "uploaded_file"))
			expectedSize := fileSize(fixturePath1)
			Expect(uploadSize).To(Equal(expectedSize))
		})

		It("overwrites the files if blobstore contains the same key", func() {
			err = fileStore.Upload("uploaded_file", fixture1)
			Expect(err).NotTo(HaveOccurred())
			err = fileStore.Upload("uploaded_file", fixture2)
			Expect(err).NotTo(HaveOccurred())

			uploadSize := fileSize(filepath.Join(blobstorePath, "uploaded_file"))
			expectedSize := fileSize(fixturePath2)
			Expect(uploadSize).To(Equal(expectedSize))
		})

		It("can accept keys in ab/cd/abcdef format", func() {
			err = fileStore.Upload("ab/cd/abcdef", fixture1)
			Expect(err).NotTo(HaveOccurred())

			uploadSize := fileSize(filepath.Join(blobstorePath, "ab/cd/abcdef"))
			expectedSize := fileSize(fixturePath1)
			Expect(uploadSize).To(Equal(expectedSize))
		})
	})
})
