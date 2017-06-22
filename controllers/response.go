package controllers

type Response struct {
	Status int         `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

func Newresponse(status int, err string, data interface{}) *Response {
	resp := &Response{
		Status: status,
		Error:  err,
		Data:   data,
	}

	return resp
}
