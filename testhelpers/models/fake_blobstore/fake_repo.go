package fake_blobstore

import (
	"github.com/cloudfoundry-incubator/api/models/blobstore"
)

type Repo struct {
	AppPackageBlobstore *Blobstore
}

func NewFakeBlobstoreRepo() *Repo {
	return new(Repo)
}

func (r *Repo) AppPackageStore() blobstore.BlobStore {
	return r.AppPackageBlobstore
}
