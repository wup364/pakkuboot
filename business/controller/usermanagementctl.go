// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// contoller 示例

package controller

import (
	"net/http"
	"pakkuboot/business/objects/cmds"
	"pakkuboot/business/service"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/serviceutil"
	"github.com/wup364/pakku/utils/strutil"
)

// UserManagementCtl 示例接口
type UserManagementCtl struct {
	usermg service.UserManagement `@autowired:""`
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
	ExecutAndAutoResponseNoConditions(r, w, func() (any, error) {
		return ctl.usermg.Query(cmds.QueryUserCmd{
			Account:  r.FormValue("account"),
			UserName: r.FormValue("userName"),
			PageableCmd: cmds.PageableCmd{
				Limit:  strutil.String2Int(r.FormValue("limit"), 100),
				Offset: strutil.String2Int(r.FormValue("offset"), 0),
			},
		})
	})
}

// createUser 查询用户
func (ctl *UserManagementCtl) createUser(w http.ResponseWriter, r *http.Request) {
	ExecutAndAutoResponse[cmds.CreateUserCmd](r, w, func(cmd cmds.CreateUserCmd) (any, error) {
		return ctl.usermg.Create(cmd)
	})
}

// updateUser 更新用户
func (ctl *UserManagementCtl) updateUser(w http.ResponseWriter, r *http.Request) {
	serviceutil.SendForbidden(w, "没有实现")
}
