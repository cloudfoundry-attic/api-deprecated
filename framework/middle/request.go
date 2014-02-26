package middle

import (
	"github.com/codegangsta/martini"
	"net/http"
)

func RequestHandler(c martini.Context, r *http.Request) {
	req := NewRequest(r)
	c.MapTo(req, (*Request)(nil))
}

type Request interface {
	Param(string) string
}

type request struct {
	req *http.Request
}

func NewRequest(req *http.Request) Request {
	return &request{req: req}
}

func (r *request) Param(key string) string {
	return r.req.FormValue(key)
}
