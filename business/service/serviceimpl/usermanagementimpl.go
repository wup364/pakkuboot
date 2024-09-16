// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 示例模块-实现UserManagement接口
package serviceimpl

import (
	"pakkuboot/business/objects/cmds"
	"pakkuboot/business/objects/dtos"
	"pakkuboot/business/repository"
	"pakkuboot/business/repository/dataobject"
	"pakkuboot/business/repository/repositoryutil"
	"pakkuboot/utils/datasource"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/logs"
)

// UserManagementImpl 示例模块
type UserManagementImpl struct {
	repository.UserRepo
	ds datasource.PakkuDataSource `@autowired:""`
}

// AsModule 作为一个模块加载
func (m *UserManagementImpl) AsModule() ipakku.Opts {
	return ipakku.Opts{
		Version:     1.0,
		Name:        "UserManagement",
		Description: "示例模块",
		OnReady: func(app ipakku.Application) {
			logs.Debugln("on ready event")
		},
		OnInit: func() {
			logs.Debugln("on init event")
		},
		OnSetup: func() {
			logs.Debugln("on setup event")
			if err := repositoryutil.ExecuteDDL(m.ds, dataobject.UserInfoPo{}); nil != err {
				logs.Panicln(err)
			}
		},
	}
}

// Create 创建用户
func (um *UserManagementImpl) Create(cmd cmds.CreateUserCmd) (res dtos.UserInfo, err error) {
	var user *dataobject.UserInfoPo
	if user, err = um.UserRepo.Create(um.ds, dataobject.UserInfoPo{
		Account:  cmd.Account,
		UserName: cmd.UserName,
		UserPWD:  cmd.UserName,
	}); nil != err {
		return
	}
	res = dtos.UserInfo{
		ID:       user.ID,
		Account:  user.Account,
		UserName: user.UserName,
		CTime:    user.CTime,
	}
	return
}

// Query 查询用户
func (um *UserManagementImpl) Query(cmd cmds.QueryUserCmd) (res dtos.PageableResult[dtos.UserInfo], err error) {
	return um.UserRepo.Query(um.ds, cmd)
}
