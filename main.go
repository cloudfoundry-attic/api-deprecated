package main

import (
	"flag"
	"fmt"
	"github.com/cloudfoundry-incubator/api/config"
	"github.com/cloudfoundry-incubator/api/router"
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

	routes := routing_table.New()

	router := router.New(router.Args{
		DefaultBackendURL: c.DefaultBackendURL,
		Routes:            routes,
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", c.Port),
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic("server exited with error: " + err.Error())
	}
}
