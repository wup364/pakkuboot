// SPDX-License-Identifier: MIT
// Copyright (C) 2023 WuPeng <wup364@outlook.com>.

package pakkudatasource

import (
	"database/sql"

	"github.com/wup364/pakku/pkg/sqlutil/sqlexecutor"
)

// PakkuDataSource 数据源
type PakkuDataSource interface {
	sqlexecutor.SqlExecutor
	TxExecutorGtter

	// GetDB 获取数据库连接
	GetDB() *sql.DB
}

// TxExecutorGtter TxExecutorGtter
type TxExecutorGtter interface {
	// GetSqlTxExecutor 开启一个事务, 返回sql执行器
	GetSqlTxExecutor() (sqlexecutor.SqlTxExecutor, error)
}
