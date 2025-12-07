// SPDX-License-Identifier: MIT
// Copyright (C) 2021 WuPeng <wup364@outlook.com>.

package pakkusys

import (
	"database/sql"

	"github.com/wup364/pakku/ipakku"
)

const (
	// C_DB_TYPE_SQLITE3 数据库类型 - sqlite3
	C_DB_TYPE_SQLITE3 = "sqlite3"

	// DEFAULT_VAL_DATAS_DIR 默认值 - 数据目录
	DEFAULT_VAL_DATAS_DIR = ".datas"

	// DEFAULT_VAL_DATASOURCE_DRIVER 默认值 - 数据源驱动
	DEFAULT_VAL_DATASOURCE_DRIVER = C_DB_TYPE_SQLITE3

	// DEFAULT_VAL_DATASOURCE_URL 默认值 - 数据源URL
	DEFAULT_VAL_DATASOURCE_URL = DEFAULT_VAL_DATAS_DIR + "/%s.db?cache=shared&_busy_timeout=60000"

	// PAKKU_PARAMS_KEY_DATASOUCE 数据源实例参数键
	PAKKU_PARAMS_KEY_DATASOUCE = "pakku.datasource.instance"
)

// PakkuApplication app实例部分接口
type PakkuApplication interface {

	// GetInstanceID 获取实例的ID
	GetInstanceID() string

	// Params 实例中的键值对数据
	Params() ipakku.Params

	// PakkuModules 默认模块Getter
	PakkuModules() ipakku.PakkuModulesGetter
}

// ModuleEvent 模块事件监听注册对象
type ModuleEvent struct {
	Module  string               // 模块名字
	Event   ipakku.ModuleEvent   // 监听事件
	Handler ipakku.OnModuleEvent // 事件处理函数
}

// OverrideModule 复写配置
type OverrideModule struct {
	Interface string      // 接口名字 ICache, Ixxx
	Implement string      // 新的接口实例注册名称
	Instance  interface{} // 可选-实现实例对象, 不为空则自动注册
}

// DataSourceInstance 数据源实例对象
type DataSourceInstance struct {
	DriverName    string `@value:"pakku.datasource.driver"`
	DataSourceUrl string `@value:"pakku.datasource.url"`
	DB            *sql.DB
}
