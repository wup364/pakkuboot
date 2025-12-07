// SPDX-License-Identifier: MIT
// Copyright (C) 2023 WuPeng <wup364@outlook.com>.

package pakkudatasource

import (
	"database/sql"
	"fmt"
	"pakkuboot/pakkusys"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/pkg/fileutil"
	"github.com/wup364/pakku/pkg/logs"
	"github.com/wup364/pakku/pkg/sqlutil/sqlexecutor"
)

// PakkuTestDataSourceImpl 测试数据源
type PakkuTestDataSourceImpl struct {
	PakkuDataSourceModule
}

// AsModule 模块加载器接口实现, 返回模块信息&配置
func (m *PakkuTestDataSourceImpl) AsModule() ipakku.Opts {
	return ipakku.Opts{
		Name:        "PakkuDataSource",
		Version:     1.0,
		Description: "测试数据源(DB)",
		OnReady: func(app ipakku.Application) {
			var err error

			// 使用sqlite3
			appName := app.Params().GetParam(ipakku.PARAMS_KEY_APPNAME).ToString(ipakku.DEFT_VAL_APPNAME)
			m.dbIns.DataSourceUrl = fmt.Sprintf(pakkusys.DEFAULT_VAL_DATASOURCE_URL, appName)
			m.dbIns.DriverName = pakkusys.C_DB_TYPE_SQLITE3

			// 如果是 sqlite3 初始化时需要创建目录
			if !fileutil.IsExist(pakkusys.DEFAULT_VAL_DATAS_DIR) {
				if err = fileutil.Mkdir(pakkusys.DEFAULT_VAL_DATAS_DIR); nil != err {
					logs.Panic(err)
					return
				}
			}

			if m.dbIns.DB, err = sql.Open(m.dbIns.DriverName, m.dbIns.DataSourceUrl); nil != err {
				logs.Panic(err)
				return
			}

			// m.dbIns = app.Params().GetParam(pakkusys.PAKKU_PARAMS_KEY_DATASOUCE).GetVal().(pakkusys.DataSourceInstance)
			m.SqlExecutor = sqlexecutor.NewSqlExecutor4Normal(m.dbIns.DriverName, m.dbIns.DB)
		},
	}
}
