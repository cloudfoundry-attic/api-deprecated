package blobstore

import (
	"io"
	"os"
	"path/filepath"
)

const (
	fileFlag = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	filePerm = 0666
	dirPerm  = 0777
)

type fileStore struct {
	path string
}

func NewFileSystemBlobStore(filepath string) BlobStore {
	return &fileStore{path: filepath}
}

func (s *fileStore) Upload(key string, content io.ReadSeeker) (err error) {
	destFilePath := filepath.Join(s.path, key)

	err = os.MkdirAll(filepath.Dir(destFilePath), dirPerm)
	if err != nil {
		return
	}

	file, err := os.OpenFile(destFilePath, fileFlag, filePerm)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	if err != nil {
		return
	}

	return
}
