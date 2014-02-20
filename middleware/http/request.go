package http

import (
	"github.com/codegangsta/martini"
	"net/http"
)

func RequestHandler(c martini.Context, req *http.Request) {
}

type Request interface {
}

type request struct {
	req *http.Request
}

func newRequest(params map[string]string, req *http.Request) Request {
	return request{
		req: req,
	}
}

type TestRequest struct {
}

func NewTestRequest() *TestRequest {
	return new(TestRequest)
}
