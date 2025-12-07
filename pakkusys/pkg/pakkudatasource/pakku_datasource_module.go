// SPDX-License-Identifier: MIT
// Copyright (C) 2023 WuPeng <wup364@outlook.com>.

package pakkudatasource

import (
	"database/sql"
	"pakkuboot/pakkusys"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/pkg/sqlutil/sqlexecutor"
)

// PakkuDataSourceModule 数据源
type PakkuDataSourceModule struct {
	sqlexecutor.SqlExecutor
	dbIns pakkusys.DataSourceInstance
}

// AsModule 模块加载器接口实现, 返回模块信息&配置
func (m *PakkuDataSourceModule) AsModule() ipakku.Opts {
	return ipakku.Opts{
		Name:        "PakkuDataSource",
		Version:     1.0,
		Description: "数据源(DB)",
		OnReady: func(app ipakku.Application) {
			m.dbIns = app.Params().GetParam(pakkusys.PAKKU_PARAMS_KEY_DATASOUCE).GetVal().(pakkusys.DataSourceInstance)
			m.SqlExecutor = sqlexecutor.NewSqlExecutor4Normal(m.dbIns.DriverName, m.dbIns.DB)
		},
	}
}

// GetDB 获取数据库连接
func (ds *PakkuDataSourceModule) GetDB() *sql.DB {
	return ds.dbIns.DB
}

// GetDriverName 获取数据库连接驱动名
func (ds *PakkuDataSourceModule) GetDriverName() string {
	return ds.dbIns.DriverName
}

// Begin 开启事务, 返回事务对象
func (ds *PakkuDataSourceModule) Begin() (*sql.Tx, error) {
	return ds.GetDB().Begin()
}

// GetSqlTxExecutor 开启一个事务, 返回sql执行器
func (ds *PakkuDataSourceModule) GetSqlTxExecutor() (sqlexecutor.SqlTxExecutor, error) {
	if tx, err := ds.Begin(); nil == err {
		return sqlexecutor.NewSqlExecutor4Tx(ds.GetDriverName(), tx), nil
	} else {
		return nil, err
	}
}
