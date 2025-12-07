// SPDX-License-Identifier: MIT
// Copyright (C) 2024 WuPeng <wup364@outlook.com>.

package dtos

// PageableResult 分页结果
type PageableResult[T any] struct {
	Total int64 `json:"total"`
	Datas []T   `json:"datas"`
}
