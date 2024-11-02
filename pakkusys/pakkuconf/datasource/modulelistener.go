// Copyright (C) 2024 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package datasource

import (
	"fmt"
	"pakkuboot/pakkusys/sysconstants"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/fileutil"
	"github.com/wup364/pakku/utils/logs"
)

// initDataSourceConfigFile 初始化数据源配置文件
func initDataSourceConfigFile(m ipakku.Modules) {
	m.OnModuleEvent(ipakku.ModuleID.AppConfig, ipakku.ModuleEventOnLoaded, func(module interface{}, app ipakku.Application) {
		//
		logs.Infoln("[模块监听] 设置默认的数据源配置文件")

		appConfig := module.(ipakku.AppConfig)
		if len(appConfig.GetConfig(sysconstants.CONFIG_KEY_DATASOURCE_DRIVER).ToString("")) > 0 {
			return // 如果配置已经设置过了, 就不继续设置了
		}

		// 设置数据源驱动配置
		if err := appConfig.SetConfig(sysconstants.CONFIG_KEY_DATASOURCE_DRIVER, sysconstants.DEFAULT_VAL_DATASOURCE_DRIVER); nil != err {
			logs.Panicln(err)
		}

		// 设置数据源URL配置
		appName := app.Params().GetParam(ipakku.PARAMS_KEY_APPNAME).ToString(ipakku.DEFT_VAL_APPNAME)
		if err := appConfig.SetConfig(sysconstants.CONFIG_KEY_DATASOURCE_URL, fmt.Sprintf(sysconstants.DEFAULT_VAL_DATASOURCE_URL, appName)); nil != err {
			logs.Panicln(err)
		}

		// 如果是 sqlite3 初始化时需要创建目录
		if sysconstants.C_DB_TYPE_SQLITE3 == sysconstants.DEFAULT_VAL_DATASOURCE_DRIVER && !fileutil.IsExist(sysconstants.DEFAULT_VAL_DATAS_DIR) {
			if err := fileutil.Mkdir(sysconstants.DEFAULT_VAL_DATAS_DIR); nil != err {
				logs.Panicln(err)
			}
		}
	})
}
