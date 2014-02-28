package blobstore_test

import (
	"bytes"
	"github.com/cloudfoundry-incubator/api/models/blobstore"
	"github.com/cloudfoundry-incubator/api/testhelpers/models/fake_s3_bucket"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("S3 Blobstore", func() {
	var (
		s3Store      blobstore.BlobStore
		fakeS3Bucket *fake_s3_bucket.FakeS3Bucket
	)
	BeforeEach(func() {
		fakeS3Bucket = &fake_s3_bucket.FakeS3Bucket{}
		s3Store = blobstore.NewS3FileSystemBlobstore(fakeS3Bucket)
	})

	Context("Upload", func() {
		It("uploads to the bucket", func() {
			content := []byte{}
			reader := bytes.NewReader(content)
			s3Store.Upload("key", reader)
			Expect(fakeS3Bucket.UploadInputs.Key).To(Equal("key"))
			Expect(fakeS3Bucket.UploadInputs.Content).To(Equal(content))
		})
	})
})
