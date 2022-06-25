// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 示例模块-实现SayHello接口
package serviceimpl

import (
	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/logs"
	"github.com/wup364/pakku/utils/strutil"
)

// SayHelloImpl 示例模块
type SayHelloImpl struct {
	conf ipakku.AppConfig `@autowired:"AppConfig"`
}

// AsModule 作为一个模块加载
func (b *SayHelloImpl) AsModule() ipakku.Opts {
	return ipakku.Opts{
		Name:        "SayHello",
		Version:     1.0,
		Description: "Say 'Hello'",
		OnReady: func(mctx ipakku.Loader) {
			logs.Debugln("on ready event")
			b.conf.SetConfig("on-ready", mctx.GetInstanceID())
		},
		OnInit: func() {
			logs.Debugln("on init event")
		},
		OnSetup: func() {
			logs.Debugln("on setup event")
			b.conf.SetConfig("on-setup", strutil.GetUUID())
		},
	}
}

func (b *SayHelloImpl) SayHello() string {
	logs.Infoln("function invoke SayHello")
	return "SayHello: " + b.conf.GetConfig("on-setup").ToString("") + ":" + b.conf.GetConfig("on-ready").ToString("") + ":" + strutil.GetUUID()
}
