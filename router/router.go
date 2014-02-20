package router

import (
	"github.com/codegangsta/martini"
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
}

func New(args Args) Router {
	url, err := url.Parse(args.DefaultBackendURL)
	if err != nil {
		panic(err)
	}

	defaultProxy := httputil.NewSingleHostReverseProxy(url)

	m := martini.Classic()
	for _, route := range args.Routes {
		switch strings.ToLower(route.Method) {
		case "get":
			m.Get(route.Path, route.Action)
		case "post":
			m.Post(route.Path, route.Action)
		case "put":
			m.Put(route.Path, route.Action)
		case "delete":
			m.Put(route.Path, route.Action)
		default:
			panic("Unknown verb: " + route.Method)
		}
	}

	m.Any("**", defaultProxy.ServeHTTP)
	return m
}
