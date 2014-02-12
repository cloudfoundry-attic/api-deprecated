package intercept_proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type Args struct {
	DefaultBackendURL string
}

type proxy struct {
	defaultBackendURL string
}

func New(args Args) Proxy {
	p := new(proxy)
	p.defaultBackendURL = args.DefaultBackendURL
	url, err := url.Parse(args.DefaultBackendURL)
	if err != nil {
		panic(err)
	}
	return httputil.NewSingleHostReverseProxy(url)
}
