// Copyright (C) 2024 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// DDL 脚本执行工具, 执行 dataobject 的脚本
package repositoryutil

import "pakkuboot/utils/datasource"

// DDLGetter DDL脚本获取
type DDLGetter interface {
	GetDDL() []string
}

// ExecuteDDL 执行dataobject里面的ddl脚本
func ExecuteDDL(ds datasource.PakkuDataSource, databaseObjects ...DDLGetter) (err error) {
	if len(databaseObjects) == 0 {
		return
	}

	// 开启事务
	var exec datasource.SqlTxExecutor
	if exec, err = ds.GetTxExecutor(); nil != err {
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
