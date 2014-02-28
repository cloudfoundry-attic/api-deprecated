package fake_s3_bucket

import (
	"io"
	"io/ioutil"
	"net/http"
)

type FakeS3Bucket struct {
	UploadInputs struct {
		Key     string
		Content []byte
	}
	UploadOutputs struct {
		Err error
	}
}

func (f *FakeS3Bucket) GetObject(key string) (data []byte, err error) {
	return
}

func (f *FakeS3Bucket) GetHeader(key string) (header http.Header, err error) {
	return
}

func (f *FakeS3Bucket) StoreObject(key string, data []byte) (err error) {
	return
}

func (f *FakeS3Bucket) DeleteObject(key string) (err error) {
	return
}

func (f *FakeS3Bucket) Put(key string, data io.ReadSeeker) (err error) {
	f.UploadInputs.Key = key
	content, readErr := ioutil.ReadAll(data)
	f.UploadInputs.Content = content
	if readErr != nil {
		panic(readErr)
	}
	return f.UploadOutputs.Err
}

func (f *FakeS3Bucket) ListKeys(prevKey string) (keys []string, err error) {
	return
}
