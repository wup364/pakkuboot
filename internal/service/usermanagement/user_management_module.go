// SPDX-License-Identifier: MIT
// Copyright (C) 2021 WuPeng <wup364@outlook.com>.

// 示例模块-实现UserManagement接口
package usermanagement

import (
	"pakkuboot/internal/objects/cmds"
	"pakkuboot/internal/objects/dtos"
	"pakkuboot/internal/repository"
	"pakkuboot/internal/repository/dataobject"
	"pakkuboot/pakkusys/pkg/pakkudatasource"
	"pakkuboot/pakkusys/pkg/repositoryutil"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/pkg/logs"
)

// UserManagementModule 示例模块
type UserManagementModule struct {
	repository.UserRepo
	ds pakkudatasource.PakkuDataSource `@autowired:""`
}

// AsModule 作为一个模块加载
func (m *UserManagementModule) AsModule() ipakku.Opts {
	return ipakku.Opts{
		Version:     1.0,
		Name:        "UserManagement",
		Description: "示例模块",
		OnReady: func(app ipakku.Application) {
			logs.Debug("on ready event")
		},
		OnInit: func() {
			logs.Debug("on init event")
		},
		OnSetup: func() {
			logs.Debug("on setup event")
			if err := repositoryutil.ExecuteDDL(m.ds, dataobject.UserInfoPo{}); nil != err {
				logs.Panic(err)
			}
		},
	}
}

// Create 创建用户
func (um *UserManagementModule) Create(cmd cmds.CreateUserCmd) (res dtos.UserInfo, err error) {
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
func (um *UserManagementModule) Query(cmd cmds.QueryUserCmd) (res dtos.PageableResult[dtos.UserInfo], err error) {
	return um.UserRepo.Query(um.ds, cmd)
}
