package blobstore_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"

	"testing"
)

func TestBlobstore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Blobstore Suite")
}

func fileSize(path string) int64 {
	stat, err := os.Stat(path)
	Expect(err).NotTo(HaveOccurred())
	return stat.Size()
}
