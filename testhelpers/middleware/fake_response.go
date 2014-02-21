package middleware

import "encoding/json"

type TestResponse struct {
	StatusCode int
	Body       string
	JsonErr    error
}

func NewTestResponse() *TestResponse {
	return new(TestResponse)
}

func (res *TestResponse) RenderJson(statusCode int, jsonObj interface{}) {
	jsonBytes, err := json.Marshal(jsonObj)
	res.Body = string(jsonBytes)
	res.StatusCode = statusCode
	res.JsonErr = err
}
