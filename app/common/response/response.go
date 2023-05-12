package response

import "encoding/json"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (res *Response) WithMessage(message string) Response {
	return Response{
		Code:    res.Code,
		Message: message,
		Data:    res.Data,
	}
}

func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code:    res.Code,
		Message: res.Message,
		Data:    data,
	}
}

func (res *Response) ToString() string {
	err := &struct {
		Code    int
		Message string
		Data    interface{}
	}{
		Code:    res.Code,
		Message: res.Message,
		Data:    res.Data,
	}

	raw, _ := json.Marshal(err)
	return string(raw)
}

func response(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

var (
	OK  = response(200, "OK")
	Err = response(500, "")
)
