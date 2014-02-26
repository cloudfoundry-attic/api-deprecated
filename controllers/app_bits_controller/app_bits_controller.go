package app_bits_controller

import (
	"github.com/cloudfoundry-incubator/api/digest"
	"github.com/cloudfoundry-incubator/api/framework/json"
	"github.com/cloudfoundry-incubator/api/framework/middle"
	"github.com/cloudfoundry-incubator/api/models/app"
	"github.com/cloudfoundry-incubator/api/models/blobstore"
	"github.com/codegangsta/martini"
	"net/http"
	"os"
)

func Put(
	req middle.Request,
	res middle.Response,
	params martini.Params,
	blobRepo blobstore.Repo,
	appRepo app.Repo,
) {

	filePath := req.Param("application_path")
	appGuid := params["app_guid"]

	file, err := os.Open(filePath)
	handleError(err)
	defer cleanupLocalFile(file)

	appModel, found := appRepo.FindByGuid(appGuid)
	if !found {
		return
	}

	appStore := blobRepo.AppPackageStore()
	err = appStore.Upload(appGuid, file)
	handleError(err)

	_, err = file.Seek(0, 0)
	handleError(err)

	digest, err := digest.Hex(file)
	handleError(err)

	appModel.SetPackageHash(digest)

	err = appRepo.Save(appModel)
	handleError(err)

	res.RenderJson(http.StatusCreated, json.Map{})
}

func cleanupLocalFile(file *os.File) {
	file.Close()
	err := os.Remove(file.Name())
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
