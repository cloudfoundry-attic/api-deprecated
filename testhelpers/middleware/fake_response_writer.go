package middleware

import "net/http"

type FakeResponseWriter struct {
	HeaderOutput http.Header
	StatusCode   int
	Body         string
}

func (res *FakeResponseWriter) Header() http.Header {
	return res.HeaderOutput
}

func (res *FakeResponseWriter) Write(data []byte) (total int, err error) {
	res.Body = string(data)
	total = len(data)
	return
}

func (res *FakeResponseWriter) WriteHeader(statusCode int) {
	res.StatusCode = statusCode
	return
}
