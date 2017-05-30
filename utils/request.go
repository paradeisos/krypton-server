package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/astaxie/beego"
)

func GetRequest(path string, params ...url.Values) *http.Request {
	var (
		request *http.Request
		err     error
	)

	if len(params) == 0 {
		request, err = http.NewRequest("GET", path, nil)
	} else {
		request, err = http.NewRequest("GET", path+"?"+params[0].Encode(), nil)
	}
	if err != nil {
		beego.Emergency(err)
	}

	return request
}

func RequestJson(method, path string, data ...interface{}) *http.Request {
	res, err := json.Marshal(data[0])
	if err != nil {
		beego.Emergency(err)
		return nil
	}

	return GenerateRequest(method, path, "application/json", string(res))
}

func RequestForm(method, path string, data ...interface{}) *http.Request {
	return GenerateRequest(method, path, "application/x-www-form-urlencoded", data)
}

func GenerateRequest(method, path, contentType string, data ...interface{}) (request *http.Request) {
	var (
		err error
	)

	if len(data) == 0 {
		request, err = http.NewRequest(method, path, nil)
	} else {
		var reader io.Reader

		body := data[0]
		switch body.(type) {
		case io.Reader:
			reader, _ = body.(io.Reader)
		case string:
			s, _ := body.(string)
			reader = bytes.NewBufferString(s)
		case url.Values:
			params, _ := body.(url.Values)
			reader = bytes.NewBufferString(params.Encode())
		}

		request, err = http.NewRequest(method, path, reader)
	}

	if err != nil {
		beego.Emergency(err)
	}

	request.Header.Set("Content-Type", contentType)

	return
}
