package blobstore

import (
	"fmt"
	"github.com/jacobsa/aws"
	"github.com/jacobsa/aws/s3"
	"io"
)

type Provider string

const (
	Local Provider = "local"
	S3    Provider = "s3"
)

type BlobStore interface {
	Upload(key string, content io.ReadSeeker) error
}

type BlobStoreArgs struct {
	Filepath        string
	Provider        Provider
	AccessKeyId     string
	AccessKeySecret string
	Host            string
	BucketName      string
}

func newBlobStore(args BlobStoreArgs) (blobstore BlobStore) {
	provider := args.Provider
	switch {
	case provider == Local:
		blobstore = NewFileSystemBlobStore(args.Filepath)
	case provider == S3:
		bucket, err := s3.OpenBucket(
			args.BucketName,
			s3.Region(args.Host),
			aws.AccessKey{
				Id:     args.AccessKeyId,
				Secret: args.AccessKeySecret,
			})

		if err != nil {
			panic(err)
		}
		blobstore = NewS3FileSystemBlobstore(bucket)
	default:
		panic(fmt.Sprintf("Empty or unknown provider [%s]\n", provider))
	}

	return
}
