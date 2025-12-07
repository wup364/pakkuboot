// SPDX-License-Identifier: MIT
// Copyright (C) 2021 WuPeng <wup364@outlook.com>.

package interfaces

import (
	"pakkuboot/internal/objects/cmds"
	"pakkuboot/internal/objects/dtos"
)

// UserManagement 示例接口
type UserManagement interface {
	// Create 创建用户
	Create(cmd cmds.CreateUserCmd) (dtos.UserInfo, error)

	//Query 查询用户
	Query(cmd cmds.QueryUserCmd) (dtos.PageableResult[dtos.UserInfo], error)
}
