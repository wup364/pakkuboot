// SPDX-License-Identifier: MIT
// Copyright (C) 2024 WuPeng <wup364@outlook.com>.

// 用户信息实体
package dataobject

// UserInfoPo 用户表存储的结构
type UserInfoPo struct {
	ID       string
	Account  string
	UserName string
	UserPWD  string
	CTime    int64
}

// GetDDL DDL脚本
func (po UserInfoPo) GetDDL() []string {
	return []string{
		`CREATE TABLE IF NOT EXISTS USER (
			ID VARCHAR(64) PRIMARY KEY,
			ACCOUNT VARCHAR(255) NOT NULL,
			USER_PWD VARCHAR(255) DEFAULT '',
			USER_NAME VARCHAR(255) NOT NULL,
			CTIME BIGINT NOT NULL
		);`,
	}
}
