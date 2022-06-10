//MIT License
//
//Copyright (c) 2022 druidcaesa
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package net

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	time int64
	url  string
	resp *http.Response
	body []byte
}

func (r *Response) Response() *http.Response {
	if r != nil {
		return r.resp
	}
	return nil
}

func (r *Response) StatusCode() int {
	if r.resp == nil {
		return 0
	}
	return r.resp.StatusCode
}

func (r *Response) Time() string {
	if r != nil {
		return fmt.Sprintf("%dms", r.time)
	}
	return "0ms"
}

func (r *Response) Url() string {
	if r != nil {
		return r.url
	}
	return ""
}

func (r *Response) Headers() http.Header {
	if r != nil {
		return r.resp.Header
	}
	return nil
}

func (r *Response) Cookies() []*http.Cookie {
	if r != nil {
		return r.resp.Cookies()
	}
	return []*http.Cookie{}
}

func (r *Response) Body() ([]byte, error) {
	if r == nil {
		return []byte{}, errors.New("gotool.HttpUtils.Response is nil.")
	}

	defer r.resp.Body.Close()

	if len(r.body) > 0 {
		return r.body, nil
	}

	if r.resp == nil || r.resp.Body == nil {
		return nil, errors.New("response or body is nil")
	}

	b, err := ioutil.ReadAll(r.resp.Body)
	if err != nil {
		return nil, err
	}
	r.body = b

	return b, nil
}

func (r *Response) Content() (string, error) {
	b, err := r.Body()
	if err != nil {
		return "", nil
	}
	return string(b), nil
}

func (r *Response) Json(v interface{}) error {
	return r.Unmarshal(v)
}

func (r *Response) Unmarshal(v interface{}) error {
	b, err := r.Body()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	return nil
}

func (r *Response) Close() error {
	if r != nil {
		return r.resp.Body.Close()
	}
	return nil
}

func (r *Response) Export() (string, error) {
	b, err := r.Body()
	if err != nil {
		return "", err
	}

	var i interface{}
	if err := json.Unmarshal(b, &i); err != nil {
		return "", errors.New("illegal json: " + err.Error())
	}

	return Export(i), nil
}
