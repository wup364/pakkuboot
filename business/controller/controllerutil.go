// Copyright (C) 2023 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package controller

import (
	"net/http"
	"pakkuboot/business/constants"

	"github.com/wup364/pakku/utils/serviceutil"
)

// ExecutAndAutoResponse 自动解析参数, 并根据有返回数据和无返回数据进行响应
func ExecutAndAutoResponse[T any](r *http.Request, w http.ResponseWriter, h any) {
	// 解析参数
	var cmd T
	if err := serviceutil.ParseHTTPRequest(r, &cmd); nil != err {
		serviceutil.SendServerError(w, err.Error())
		return
	}

	// 执行
	if fh, ok := h.(func(h T) (any, error)); ok {
		if res, err := fh(cmd); nil != err {
			serviceutil.SendBusinessError(w, err)
		} else {
			serviceutil.SendSuccess(w, res)
		}

	} else if fh, ok := h.(func(h T) error); ok {
		if err := fh(cmd); nil != err {
			serviceutil.SendBusinessError(w, err)
		} else {
			serviceutil.SendSuccess(w, "")
		}
	} else {
		serviceutil.SendServerError(w, constants.ErrUnknownError.Error())
	}
}

// ExecutAndAutoResponseNoConditions 没有请求参数, 根据有返回数据和无返回数据进行响应
func ExecutAndAutoResponseNoConditions(r *http.Request, w http.ResponseWriter, h any) {
	// 执行
	if fh, ok := h.(func() (any, error)); ok {
		if res, err := fh(); nil != err {
			serviceutil.SendBusinessError(w, err)
		} else {
			serviceutil.SendSuccess(w, res)
		}

	} else if fh, ok := h.(func() error); ok {
		if err := fh(); nil != err {
			serviceutil.SendBusinessError(w, err)
		} else {
			serviceutil.SendSuccess(w, "")
		}
	} else {
		serviceutil.SendServerError(w, constants.ErrUnknownError.Error())
	}
}

// ParseHTTPRequest 转换request里面的json为对象, 如果出错, 则响应错误
func ParseHTTPRequest(r *http.Request, w http.ResponseWriter, cmd any) (err error) {
	if err = serviceutil.ParseHTTPRequest(r, cmd); nil != err {
		serviceutil.SendServerError(w, err.Error())
	}
	return
}
