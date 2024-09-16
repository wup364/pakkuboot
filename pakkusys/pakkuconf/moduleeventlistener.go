// Copyright (C) 2024 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package pakkuconf

import (
	"fmt"
	"pakkuboot/pakkusys"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/fileutil"
	"github.com/wup364/pakku/utils/logs"
)

// initDefaultDataSource 设置默认的数据源配置
func initDefaultDataSource() pakkusys.ModuleEvent {
	return pakkusys.ModuleEvent{
		Module: ipakku.ModuleID.AppConfig,
		Event:  ipakku.ModuleEventOnSetupSucced,
		Handler: func(module interface{}, app ipakku.Application) {
			logs.Infoln("[模块监听] 设置默认的数据源配置, 若无需要请删除")

			app.Modules().OnModuleEvent(ipakku.ModuleID.AppConfig, ipakku.ModuleEventOnLoaded, func(module interface{}, app ipakku.Application) {
				appConfig := module.(ipakku.AppConfig)
				if len(appConfig.GetConfig("datasource.driver").ToString("")) > 0 {
					return
				}

				if err := appConfig.SetConfig("datasource.driver", "sqlite3"); nil != err {
					logs.Panicln(err)
				}

				appName := app.Params().GetParam(ipakku.PARAMS_KEY_APPNAME).ToString(ipakku.DEFT_VAL_APPNAME)
				if err := appConfig.SetConfig("datasource.url", fmt.Sprintf(".datas/%s.db?cache=shared", appName)); nil != err {
					logs.Panicln(err)
				}

				if !fileutil.IsExist(".datas") {
					if err := fileutil.Mkdir(".datas"); nil != err {
						logs.Panicln(err)
					}
				}
			})
		},
	}
}
