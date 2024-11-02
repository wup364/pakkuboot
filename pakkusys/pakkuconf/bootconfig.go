// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 程序启动前的相关配置
package pakkuconf

import (
	"io"
	"pakkuboot/business/controller"
	"pakkuboot/business/service/serviceimpl"
	"pakkuboot/pakkusys"
	"pakkuboot/pakkusys/pakkuconf/datasource"
	"pakkuboot/pakkusys/pakkuconf/logger"
	"pakkuboot/pakkusys/sysmodules/pakkudatasource"

	"github.com/wup364/pakku/ipakku"
)

// EnablePakkuModules 加载额外的自带模块
func EnablePakkuModules(modules ipakku.PakkuModuleBuilder) {
	modules.EnableAppEvent()
}

// RegisterModules 注册需要加载的模块
func RegisterModules(art pakkusys.ApplicationRT) []ipakku.Module {
	return []ipakku.Module{
		new(pakkudatasource.PakkuDataSourceImpl),
		new(serviceimpl.UserManagementImpl),
	}
}

// RegisterHttpController 注册需要加载的controller
func RegisterHttpController(art pakkusys.ApplicationRT) []ipakku.Controller {
	return []ipakku.Controller{
		new(controller.UserManagementCtl),
	}
}

// RegisterHttpRequestFilter 注册需要加载的http过滤器
func RegisterHttpRequestFilter(art pakkusys.ApplicationRT) []ipakku.FilterConfigItem {
	return []ipakku.FilterConfigItem{}
}

// RegisterRPCService 注册需要加载的gorpc
func RegisterRPCService(art pakkusys.ApplicationRT) []interface{} {
	return []interface{}{}
}

// RegisterModuleEvent 注册模块加载事件
func RegisterModuleEvent(art pakkusys.ApplicationRT) []pakkusys.ModuleEvent {
	return []pakkusys.ModuleEvent{
		datasource.OnAppConfigSetupSucced(),
	}
}

// RegisterOverride 注册复写的模块
func RegisterOverride(art pakkusys.ApplicationRT) []pakkusys.OverrideModule {
	return []pakkusys.OverrideModule{
		// {
		// 	Interface: "ICache",
		// 	Implement: "redis",
		// 	Instance:  nil,
		// },
	}
}

// RegisterLoggerWriter 设置日志持久化写入器
func RegisterLoggerWriter(logdir, logName string) io.Writer {
	return logger.NewLoggerWriter4File(logdir, logName)
}
