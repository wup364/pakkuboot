// SPDX-License-Identifier: MIT
// Copyright (C) 2023 WuPeng <wup364@outlook.com>.

package cmds

// PageableCmd 分页参数
type PageableCmd struct {
	Limit  int `json:"limit" form:"limit"`
	Offset int `json:"offset" form:"offset"`
}
