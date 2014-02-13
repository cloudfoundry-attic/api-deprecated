package main

import (
	"flag"
	"fmt"
	"github.com/cloudfoundry-incubator/api/config"
	"github.com/cloudfoundry-incubator/api/intercept_proxy"
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

	proxy := intercept_proxy.New(intercept_proxy.Args{
		DefaultBackendURL: c.DefaultBackendURL,
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", c.Port),
		Handler: proxy,
	}

	fmt.Printf("%#v", c)
	err = server.ListenAndServe()
	if err != nil {
		panic("server exited with error: " + err.Error())
	}
}
