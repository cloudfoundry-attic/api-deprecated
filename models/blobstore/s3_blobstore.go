package blobstore

import (
	"github.com/jacobsa/aws/s3"
	"io"
)

type s3store struct {
	bucket s3.Bucket
}

func NewS3FileSystemBlobstore(bucket s3.Bucket) BlobStore {
	return &s3store{bucket: bucket}
}

func (s *s3store) Upload(key string, content io.ReadSeeker) (err error) {
	return s.bucket.Put(key, content)
}
