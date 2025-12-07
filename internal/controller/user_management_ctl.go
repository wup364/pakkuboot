// SPDX-License-Identifier: MIT
// Copyright (C) 2021 WuPeng <wup364@outlook.com>.

// contoller 示例

package controller

import (
	"net/http"
	"pakkuboot/internal/interfaces"
	"pakkuboot/internal/objects/cmds"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/pkg/serviceutil"
)

// UserManagementCtl 示例接口
type UserManagementCtl struct {
	usermg interfaces.UserManagement `@autowired:""`
}

// AsController 实现 AsController 接口
func (ctl *UserManagementCtl) AsController() ipakku.ControllerConfig {
	return ipakku.ControllerConfig{
		RequestMapping: "/user/v1",
		RouterConfig: ipakku.RouterConfig{
			ToLowerCase: true,
			HandlerFunc: [][]interface{}{
				{http.MethodGet, "", ctl.queryUser},
				{http.MethodPost, "", ctl.createUser},
				{http.MethodPut, ":*", ctl.updateUser},
			},
		},
		FilterConfig: []ipakku.FilterConfigItem{
			{
				Path: "",
				Func: ipakku.Filter4Passed,
			},
		},
	}
}

// queryUser 查询用户
func (ctl *UserManagementCtl) queryUser(w http.ResponseWriter, r *http.Request) {
	serviceutil.HandleRequest[cmds.QueryUserCmd](r, w, func(cmd cmds.QueryUserCmd) (any, error) {
		return ctl.usermg.Query(cmd)
	})
}

// createUser 查询用户
func (ctl *UserManagementCtl) createUser(w http.ResponseWriter, r *http.Request) {
	serviceutil.HandleRequest[cmds.CreateUserCmd](r, w, func(cmd cmds.CreateUserCmd) (any, error) {
		return ctl.usermg.Create(cmd)
	})
}

// updateUser 更新用户
func (ctl *UserManagementCtl) updateUser(w http.ResponseWriter, r *http.Request) {
	serviceutil.SendForbidden(w, "没有实现")
}
