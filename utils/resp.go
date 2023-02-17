package utils

type Response struct {
	Code     int               `json:"code"`
	Msg      string            `json:"msg"`
	Data     interface{}       `json:"data"`
	Redirect map[string]string `json:"redirect"`
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
		Code:     SuccCode,
		Msg:      SuccMsg,
		Redirect: map[string]string{"url": url},
	}
}
