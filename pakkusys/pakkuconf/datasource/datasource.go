// Copyright (C) 2024 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 数据源配置
package datasource

import (
	"database/sql"
	"pakkuboot/pakkusys"
	"pakkuboot/pakkusys/sysconstants"
	"time"

	"github.com/wup364/pakku/ipakku"
)

// OnAppConfigSetupSucced 配置模块成功安装结束事件监听
func OnAppConfigSetupSucced() pakkusys.ModuleEvent {
	return pakkusys.ModuleEvent{
		Module: ipakku.ModuleID.AppConfig,
		Event:  ipakku.ModuleEventOnSetupSucced,
		Handler: func(module interface{}, app ipakku.Application) {
			// 初始化数据源配置文件
			initDataSourceConfigFile(app.Modules())
		},
	}
}

// GetDataSource 初始化数据源配置, 获得sql.DB对象
func GetDataSource(dsSetting DataSourceSetting) (db *sql.DB, err error) {
	if db, err = sql.Open(dsSetting.DriverName, dsSetting.DataSourceName); nil == err {
		if dsSetting.DriverName == sysconstants.C_DB_TYPE_SQLITE3 {
			// database is locked
			// https://github.com/mattn/go-sqlite3/issues/209
			// ds.db.SetMaxOpenConns(1)
		} else {
			// ds.db.SetMaxIdleConns(250)
			db.SetConnMaxLifetime(time.Hour)
		}
	}
	return
}
