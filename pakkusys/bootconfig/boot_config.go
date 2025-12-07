// SPDX-License-Identifier: MIT
// Copyright (C) 2021 WuPeng <wup364@outlook.com>.

// 应用启动配置 - 加载模块、注册控制器、设置日志等
package bootconfig

import (
	"database/sql"
	"fmt"
	"io"
	"pakkuboot/internal/controller"
	"pakkuboot/internal/service/usermanagement"
	"pakkuboot/pakkusys"
	"pakkuboot/pakkusys/pkg/logger"
	"pakkuboot/pakkusys/pkg/pakkudatasource"
	"time"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/pkg/fileutil"
)

// EnablePakkuModules 加载额外的自带模块
func EnablePakkuModules(modules ipakku.PakkuModuleBuilder) {
	modules.EnableAppEvent()
}

// RegisterModules 注册需要加载的模块
func RegisterModules(art pakkusys.PakkuApplication) []ipakku.Module {
	return []ipakku.Module{
		new(pakkudatasource.PakkuDataSourceModule),
		new(usermanagement.UserManagementModule),
	}
}

// RegisterHttpController 注册需要加载的controller
func RegisterHttpController(art pakkusys.PakkuApplication) []ipakku.Controller {
	return []ipakku.Controller{
		new(controller.UserManagementCtl),
	}
}

// RegisterHttpRequestFilter 注册需要加载的http过滤器
func RegisterHttpRequestFilter(art pakkusys.PakkuApplication) []ipakku.FilterConfigItem {
	return []ipakku.FilterConfigItem{}
}

// RegisterRPCService 注册需要加载的gorpc
func RegisterRPCService(art pakkusys.PakkuApplication) []any {
	return []any{}
}

// RegisterModuleEvent 注册模块加载事件
func RegisterModuleEvent(art pakkusys.PakkuApplication) []pakkusys.ModuleEvent {
	return []pakkusys.ModuleEvent{}
}

// RegisterOverride 注册复写的模块
func RegisterOverride(art pakkusys.PakkuApplication) []pakkusys.OverrideModule {
	return []pakkusys.OverrideModule{
		// {
		// 	Interface: "ICache",
		// 	Implement: "redis",
		// 	Instance:  nil,
		// },
	}
}

// RegisterDataSource 初始化数据源配置, 获得sql.DB对象
// 在配置模块初始完成之后执行
func RegisterDataSource(art pakkusys.PakkuApplication, config ipakku.AppConfig) (dbIns *pakkusys.DataSourceInstance, err error) {
	dbIns = &pakkusys.DataSourceInstance{}
	if err = config.ScanAndAutoValue("", dbIns); nil != err {
		return
	}

	// 如果没有设置数据源, 则默认使用sqlite3
	if dbIns.DriverName == "" || dbIns.DataSourceUrl == "" {
		appName := art.Params().GetParam(ipakku.PARAMS_KEY_APPNAME).ToString(ipakku.DEFT_VAL_APPNAME)
		dbIns.DataSourceUrl = fmt.Sprintf(pakkusys.DEFAULT_VAL_DATASOURCE_URL, appName)
		dbIns.DriverName = pakkusys.C_DB_TYPE_SQLITE3

		// 如果是 sqlite3 初始化时需要创建目录
		if !fileutil.IsExist(pakkusys.DEFAULT_VAL_DATAS_DIR) {
			if err = fileutil.Mkdir(pakkusys.DEFAULT_VAL_DATAS_DIR); nil != err {
				return
			}
		}
	}

	if dbIns.DB, err = sql.Open(dbIns.DriverName, dbIns.DataSourceUrl); nil == err {
		if dbIns.DriverName == pakkusys.C_DB_TYPE_SQLITE3 {
			// database is locked
			// https://github.com/mattn/go-sqlite3/issues/209
			// ds.db.SetMaxOpenConns(1)
		} else {
			// ds.db.SetMaxIdleConns(250)
			dbIns.DB.SetConnMaxLifetime(time.Hour)
		}
	}

	return
}

// RegisterLoggerWriter 设置日志持久化写入器
func RegisterLoggerWriter(logdir, logName string, maxLogFiles int) io.Writer {
	return logger.NewLoggerWriter4File(logdir, logName, maxLogFiles)
}
