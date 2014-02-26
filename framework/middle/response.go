package middle

import (
	"encoding/json"
	"github.com/codegangsta/martini"
	"net/http"
)

func ResponseHandler(c martini.Context, w http.ResponseWriter) {
	res := NewResponse(w)
	c.MapTo(res, (*Response)(nil))
}

type Response interface {
	RenderJson(statusCode int, jsonObject interface{})
}

type response struct {
	writer     http.ResponseWriter
	statusCode int
}

func NewResponse(writer http.ResponseWriter) Response {
	return &response{
		writer: writer,
	}
}

func (res *response) RenderJson(statusCode int, jsonObject interface{}) {
	jsonBytes, err := json.MarshalIndent(jsonObject, "", "  ")
	if err != nil {
		panic(err)
	}
	res.writer.WriteHeader(statusCode)
	res.writer.Write(jsonBytes)
}
