// Copyright (C) 2024 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 用户信息维护
package repository

import (
	"database/sql"
	"pakkuboot/business/objects/cmds"
	"pakkuboot/business/objects/dtos"
	"pakkuboot/business/repository/dataobject"
	"time"

	"github.com/wup364/pakku/utils/constants/sqlconditions"
	"github.com/wup364/pakku/utils/sqlexecutor"
	"github.com/wup364/pakku/utils/sqlutil"
	"github.com/wup364/pakku/utils/strutil"
)

const (
	// sql_user_create 用户创建
	sql_user_create = `INSERT INTO USER(ID, ACCOUNT, USER_NAME, USER_PWD, CTIME) VALUES(?, ?, ?, ?, ?)`

	// sql_user_qry query
	sql_user_qry = `SELECT ID, ACCOUNT, USER_NAME, CTIME FROM USER`

	// sql_user_count count
	sql_user_count = `SELECT COUNT(ID) FROM USER`
)

// UserRepo 用户信息维护
type UserRepo struct{}

// Create 新建用户
func (repo *UserRepo) Create(exec sqlexecutor.Exec, user dataobject.UserInfoPo) (res *dataobject.UserInfoPo, err error) {
	if len(user.ID) == 0 {
		user.ID = strutil.GetUUID()
	}
	user.CTime = time.Now().UnixMilli()
	if _, err = exec.ExecWithPrepare(sql_user_create, user.ID, user.Account, user.UserName, repo.encodePwd(user.UserPWD), user.CTime); nil != err {
		return
	} else {
		res = &user
	}
	return
}

// Query 查询用户
func (repo *UserRepo) Query(exec sqlexecutor.Query, cmd cmds.QueryUserCmd) (res dtos.PageableResult[dtos.UserInfo], err error) {
	//
	if res.Total, err = repo.CountQuery(exec, cmd); nil != err || res.Total == 0 {
		return
	}

	//
	conditions := []string{"ACCOUNT = ?", "USER_NAME like ?"}
	args := strutil.ToInterface(strutil.RemoveEmpty(cmd.Account, cmd.UserName)...)

	// 执行count查询
	var rows *sql.Rows
	sqlstr := sqlutil.SqlConditionConcatForWhere(sql_user_qry, sqlconditions.AND, conditions, cmd.Account, cmd.UserName)
	if rows, err = exec.QueryWithPrepare(sqlstr, args...); nil != err {
		sqlutil.CloseRowsSilence(rows)
		return
	}

	res.Datas, err = sqlutil.ScanAndClose[dtos.UserInfo](rows, func(scan func(...any) error) (obj dtos.UserInfo, err error) {
		return obj, scan(&obj.ID, &obj.Account, &obj.UserName, &obj.CTime)
	})

	return
}

// CountQuery count查询用户
func (repo *UserRepo) CountQuery(exec sqlexecutor.Query, cmd cmds.QueryUserCmd) (res int64, err error) {
	//
	conditions := []string{"ACCOUNT = ?", "USER_NAME like ?"}
	args := strutil.ToInterface(strutil.RemoveEmpty(cmd.Account, cmd.UserName)...)

	// 执行count查询
	var rows *sql.Rows
	sqlstr := sqlutil.SqlConditionConcatForWhere(sql_user_count, sqlconditions.AND, conditions, cmd.Account, cmd.UserName)
	if rows, err = exec.QueryWithPrepare(sqlstr, args...); nil != err {
		sqlutil.CloseRowsSilence(rows)
		return
	}

	return sqlutil.ScanFirstOneAndClose[int64](rows)
}

// encodePwd 加密密码
func (repo *UserRepo) encodePwd(pwd string) string {
	return strutil.GetSHA256(pwd)
}
