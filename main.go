package main

import (
	"flag"
	"fmt"
	"github.com/cloudfoundry-incubator/api/config"
	"github.com/cloudfoundry-incubator/api/framework/database"
	"github.com/cloudfoundry-incubator/api/framework/router"
	"github.com/cloudfoundry-incubator/api/models"
	"github.com/cloudfoundry-incubator/api/routing_table"
	"net/http"
)

var configPath *string

func init() {
	configPath = flag.String("c", "config.yml", "path to config file")
	flag.Parse()
}

func main() {
	c, err := config.NewFromFile(*configPath)
	if err != nil {
		panic("error reading config file: " + err.Error())
	}

	db, err := database.NewDB(c.DB)
	if err != nil {
		panic("error connecting to db: " + err.Error())
	}

	router := router.New(router.Args{
		DefaultBackendURL: c.DefaultBackendURL,
		Routes:            routing_table.New(),
		Dependencies:      models.NewExports(db, c),
	})

	address := fmt.Sprintf(":%d", c.Port)

	server := http.Server{
		Addr:    address,
		Handler: router,
	}

	fmt.Printf("Cloud Foundry api server listening on %s\n", address)

	err = server.ListenAndServe()
	if err != nil {
		panic("server exited with error: " + err.Error())
	}
}
