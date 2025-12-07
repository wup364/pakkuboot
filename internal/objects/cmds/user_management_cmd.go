// SPDX-License-Identifier: MIT
// Copyright (C) 2024 WuPeng <wup364@outlook.com>.

package cmds

// CreateUserCmd 创建用户
type CreateUserCmd struct {
	Account  string
	UserName string
	Passwd   string
}

// QueryUserCmd 查询用户
type QueryUserCmd struct {
	PageableCmd
	Account  string `form:"account"`
	UserName string `ffmr:"userName"`
}
