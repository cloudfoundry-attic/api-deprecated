package blobstore

type Repo interface {
	AppPackageStore() BlobStore
}

type RepoArgs struct {
	AppPackageStore BlobStoreArgs
}

type repo struct {
	appPackageStore BlobStore
}

func NewRepo(args RepoArgs) Repo {
	return &repo{
		appPackageStore: newBlobStore(args.AppPackageStore),
	}
}

func (r *repo) AppPackageStore() (store BlobStore) {
	return r.appPackageStore
}
