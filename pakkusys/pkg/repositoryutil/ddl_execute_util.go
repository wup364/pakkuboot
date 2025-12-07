// SPDX-License-Identifier: MIT
// Copyright (C) 2024 WuPeng <wup364@outlook.com>.

// DDL 脚本执行工具, 执行 dataobject 的脚本
package repositoryutil

import (
	"pakkuboot/pakkusys/pkg/pakkudatasource"

	"github.com/wup364/pakku/pkg/sqlutil/sqlexecutor"
)

// DDLGetter DDL脚本获取
type DDLGetter interface {
	GetDDL() []string
}

// ExecuteDDL 执行dataobject里面的ddl脚本
func ExecuteDDL(ds pakkudatasource.TxExecutorGtter, databaseObjects ...DDLGetter) (err error) {
	if len(databaseObjects) == 0 {
		return
	}

	// 开启事务
	var exec sqlexecutor.SqlTxExecutor
	if exec, err = ds.GetSqlTxExecutor(); nil != err {
		return
	}

	// ddl exec
	for _, po := range databaseObjects {
		if ddls := po.GetDDL(); len(ddls) > 0 {
			for i := 0; i < len(ddls); i++ {
				if _, err = exec.Exec(ddls[i]); nil != err {
					exec.RollbackSilence()
					return
				}
			}
		}
	}
	return exec.Complete()
}
