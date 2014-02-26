package router

import (
	"github.com/cloudfoundry-incubator/api/framework/middle"
	"github.com/codegangsta/martini"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Route struct {
	Method string
	Path   string
	Action interface{}
}

type Router interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type Args struct {
	DefaultBackendURL string
	Routes            []Route
	Dependencies      map[interface{}]interface{}
}

func New(args Args) Router {
	url, err := url.Parse(args.DefaultBackendURL)
	if err != nil {
		panic(err)
	}

	defaultProxy := httputil.NewSingleHostReverseProxy(url)

	m := martini.Classic()
	m.Use(middle.RequestHandler)
	m.Use(middle.ResponseHandler)

	for interfaceType, dependency := range args.Dependencies {
		m.MapTo(dependency, interfaceType)
	}

	for _, route := range args.Routes {
		switch strings.ToLower(route.Method) {
		case "get":
			m.Get(route.Path, route.Action)
		case "post":
			m.Post(route.Path, route.Action)
		case "put":
			m.Put(route.Path, route.Action)
		case "delete":
			m.Delete(route.Path, route.Action)
		default:
			panic("Unknown verb: " + route.Method)
		}
	}

	m.Any("**", defaultProxy.ServeHTTP)
	return m
}
