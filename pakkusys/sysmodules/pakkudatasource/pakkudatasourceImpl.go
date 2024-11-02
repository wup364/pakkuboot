// Copyright (C) 2023 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package pakkudatasource

import (
	"database/sql"
	"pakkuboot/pakkusys/pakkuconf/datasource"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/logs"
	"github.com/wup364/pakku/utils/sqlexecutor"
)

// PakkuDataSourceImpl 数据源
type PakkuDataSourceImpl struct {
	sqlexecutor.SqlExecutor
	db        *sql.DB
	dsSetting *datasource.DataSourceSetting `@autoConfig:""`
}

// AsModule 模块加载器接口实现, 返回模块信息&配置
func (m *PakkuDataSourceImpl) AsModule() ipakku.Opts {
	return ipakku.Opts{
		Name:        "PakkuDataSource",
		Version:     1.0,
		Description: "数据源(DB)",
		OnInit: func() {
			if err := m.initial(); nil != err {
				logs.Panicln(err)
			}
		},
	}
}

// GetDB 获取数据库连接
func (ds *PakkuDataSourceImpl) GetDB() *sql.DB {
	return ds.db
}

// GetDriverName 获取数据库连接驱动名
func (ds *PakkuDataSourceImpl) GetDriverName() string {
	return ds.dsSetting.DriverName
}

// Begin 开启事务, 返回事务对象
func (ds *PakkuDataSourceImpl) Begin() (*sql.Tx, error) {
	return ds.db.Begin()
}

// GetTxExecutor 开启一个事务, 返回sql执行器
func (ds *PakkuDataSourceImpl) GetTxExecutor() (sqlexecutor.SqlTxExecutor, error) {
	if tx, err := ds.Begin(); nil == err {
		return sqlexecutor.NewSqlExecutor4Tx(ds.GetDriverName(), tx), nil
	} else {
		return nil, err
	}
}

// initial 初始化数据源配置
func (ds *PakkuDataSourceImpl) initial() (err error) {
	if nil != ds.db {
		return
	}

	if ds.db, err = datasource.GetDataSource(*ds.dsSetting); nil == err {
		// 初始化其他对象
		ds.SqlExecutor = sqlexecutor.NewSqlExecutor4Normal(ds.dsSetting.DriverName, ds.db)
	}
	return
}
