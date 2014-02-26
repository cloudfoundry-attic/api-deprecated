package middleware

type TestRequest struct {
	Params map[string]string
}

func NewTestRequest() *TestRequest {
	return new(TestRequest)
}

func (req *TestRequest) Param(key string) (value string) {
	return req.Params[key]
}
