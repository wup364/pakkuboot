// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// contoller 示例

package controller

import (
	"net/http"
	"pakkuboot/business/service"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/serviceutil"
)

// SayHelloCtl 示例接口
type SayHelloCtl struct {
	sayHello service.SayHello `@autowired:"SayHello"`
}

// AsController 实现 AsController 接口
func (ctl *SayHelloCtl) AsController() ipakku.ControllerConfig {
	return ipakku.ControllerConfig{
		RequestMapping: "/sayhello/v1",
		RouterConfig: ipakku.RouterConfig{
			ToLowerCase: true,
			HandlerFunc: [][]interface{}{
				{http.MethodGet, ctl.SayHello},
				{http.MethodGet, "/hello", ctl.SayHello},
			},
		},
		FilterConfig: []ipakku.FilterConfigItem{
			{
				Path: "",
				Func: ipakku.Filter4Passed,
			},
		},
	}
}

// 示例接口
func (ctl *SayHelloCtl) SayHello(w http.ResponseWriter, r *http.Request) {
	serviceutil.SendSuccess(w, ctl.sayHello.SayHello())
}
