package http

import (
	"encoding/json"
)

type Response interface {
	WriteJson(interface{})
	SetStatusCode(int)
}

type TestResponse struct {
	StatusCode int
	Body       string
	JsonErr    error
}

func NewTestResponse() *TestResponse {
	return new(TestResponse)
}

func (res *TestResponse) WriteJson(jsonObj interface{}) {
	jsonBytes, err := json.Marshal(jsonObj)
	res.Body = string(jsonBytes)
	res.JsonErr = err
}

func (res *TestResponse) SetStatusCode(code int) {
	res.StatusCode = code
}
