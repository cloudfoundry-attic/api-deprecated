package blobstore

import (
	"io"
)

type BlobStore interface {
	Upload(key string, content io.ReadSeeker) error
}

type BlobStoreArgs struct {
	Filepath string
}

func NewBlobStore(args BlobStoreArgs) BlobStore {
	return newFileSystemBlobStore(args.Filepath)
}
