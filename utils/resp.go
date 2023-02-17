package utils

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Url  string      `json:"url"`
}

func Succ(data interface{}) *Response {
	return &Response{
		Code: SuccCode,
		Msg:  SuccMsg,
		Data: data,
	}
}

func Fail(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func SuccWithUrl(url string) *Response {
	return &Response{
		Code: SuccCode,
		Msg:  SuccMsg,
		Data: data,
	}
}