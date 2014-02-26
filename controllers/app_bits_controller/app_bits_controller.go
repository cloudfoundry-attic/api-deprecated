package app_bits_controller

import (
	"github.com/cloudfoundry-incubator/api/framework/json"
	"github.com/cloudfoundry-incubator/api/framework/middle"
	"github.com/cloudfoundry-incubator/api/models/blobstore"
	"github.com/codegangsta/martini"
	"net/http"
	"os"
)

func Put(
	req middle.Request,
	res middle.Response,
	params martini.Params,
	repo blobstore.Repo,
) {

	filePath := req.Param("application_path")
	appGuid := params["app_guid"]

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	appStore := repo.AppPackageStore()
	err = appStore.Upload(appGuid, file)
	if err != nil {
		panic(err)
	}

	res.RenderJson(http.StatusOK, json.Map{})
}
