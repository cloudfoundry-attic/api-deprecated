package fake_blobstore

import (
	"io"
	"io/ioutil"
)

type Blobstore struct {
	UploadInputs struct {
		Key  string
		File []byte
	}
	UploadOutputs struct {
		Err error
	}
}

func NewFakeBlobstore() *Blobstore {
	return new(Blobstore)
}

func (b *Blobstore) Upload(key string, r io.ReadSeeker) (err error) {
	b.UploadInputs.Key = key
	uploadedFile, readErr := ioutil.ReadAll(r)
	b.UploadInputs.File = uploadedFile
	if readErr != nil {
		panic(readErr)
	}
	err = b.UploadOutputs.Err
	return
}
