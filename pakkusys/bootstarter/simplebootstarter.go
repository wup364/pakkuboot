// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 实例化一个pakku运行环境

package bootstarter

import (
	"pakkuboot/pakkusys/pakkuconf"
	"pakkuboot/pakkusys/sysconstants"
	"path/filepath"
	"sync"

	"github.com/wup364/pakku"
	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/fileutil"
	"github.com/wup364/pakku/utils/logs"
)

// NewSimpleBootStarter 新建应用
func NewSimpleBootStarter(name string) *SimpleBootStarter {
	return &SimpleBootStarter{locker: new(sync.Mutex), builder: pakku.NewApplication(name)}
}

// SimpleBootStarter 简单启动器
type SimpleBootStarter struct {
	locker   *sync.Mutex
	pakkuapp ipakku.PakkuApplication
	builder  ipakku.ApplicationBootBuilder
}

// ApplicationBootBuilder 获取应用构建器
func (boot *SimpleBootStarter) ApplicationBootBuilder() ipakku.ApplicationBootBuilder {
	return boot.builder
}

// SetLogger 初始化默认的本地日志
func (boot *SimpleBootStarter) SetLogger(logger, logdir, loglevel string) *SimpleBootStarter {
	if logger == "file" {
		logdir = initialLogdir(logdir)
		appParams := boot.builder.Application().Params()
		instanceID := boot.builder.Application().GetInstanceID()
		logName := appParams.GetParam(ipakku.PARAMS_KEY_APPNAME).ToString(instanceID)
		loggerWriter := pakkuconf.RegisterLoggerWriter(logdir, logName)

		logs.Infoln("Logger output file, path: ", logdir)
		boot.builder.PakkuConfigure().SetLoggerOutput(loggerWriter)
	}

	switch loglevel {
	case "none":
		boot.builder.PakkuConfigure().SetLoggerLevel(logs.NONE)
	case "error":
		boot.builder.PakkuConfigure().SetLoggerLevel(logs.ERROR)
	case "info":
		boot.builder.PakkuConfigure().SetLoggerLevel(logs.INFO)
	default:
		boot.builder.PakkuConfigure().SetLoggerLevel(logs.DEBUG)
	}
	return boot
}

// BootStart 加载&启动程序
func (boot *SimpleBootStarter) BootStart() ipakku.PakkuApplication {
	boot.locker.Lock()
	defer boot.locker.Unlock()

	if boot.pakkuapp == nil {
		app := boot.builder.Application()
		// 加载额外的自带模块
		pakkuconf.EnablePakkuModules(boot.builder.PakkuModules())
		// 加载自定义模块
		if modules := pakkuconf.RegisterModules(app); len(modules) > 0 {
			boot.builder.CustomModules().AddModules(modules...)
		}

		// 注册覆盖模块
		if overrides := pakkuconf.RegisterOverride(app); len(overrides) > 0 {
			for i := 0; i < len(overrides); i++ {
				if overrides[i].Instance != nil {
					ipakku.PakkuConf.RegisterPakkuModuleImplement(overrides[i].Instance, overrides[i].Interface, overrides[i].Implement)
				}
				ipakku.PakkuConf.SetPakkuModuleImplement(app.Params(), overrides[i].Interface, overrides[i].Implement)
			}
		}

		// 注册模块加载事件
		if events := pakkuconf.RegisterModuleEvent(app); len(events) > 0 {
			for i := 0; i < len(events); i++ {
				boot.builder.ModuleEvents().Listen(events[i].Module, events[i].Event, events[i].Handler)
			}
		}

		boot.pakkuapp = boot.builder.BootStart()
	}

	return boot.pakkuapp
}

// BootStartWeb 以web服务方式启动, https证书放在conf目录下自动加载
// 自动装载模块: AppConfig AppService
func (boot *SimpleBootStarter) BootStartWeb(debug bool) {
	// bootstart
	if boot.pakkuapp == nil {
		boot.builder.PakkuModules().EnableAppConfig().EnableAppService()
		boot.BootStart()
	}

	// 注册controller
	service := boot.pakkuapp.PakkuModules().GetAppService()
	for _, v := range pakkuconf.RegisterHttpController(boot.pakkuapp) {
		if err := service.AsController(v); nil != err {
			logs.Panicln(err)
		}
	}

	// 注册http过滤器
	for _, v := range pakkuconf.RegisterHttpRequestFilter(boot.pakkuapp) {
		if err := service.Filter(v.Path, v.Func); nil != err {
			logs.Panicln(err)
		}
	}

	// 读取配置 & 启动服务
	certFile, keyFile := getCertFile()
	config := boot.pakkuapp.PakkuModules().GetAppConfig()
	service.StartHTTP(ipakku.HTTPServiceConfig{
		Debug:      debug,
		ListenAddr: config.GetConfig(sysconstants.CONFIG_KEY_HTTP_LISTEN_ADDRESS).ToString(sysconstants.DEFAULT_VAL_LISTEN_ADDRESS),
		CertFile:   certFile,
		KeyFile:    keyFile,
	})
}

// BootStartRpc 以Rpc服务方式启动
// 自动装载模块: AppConfig AppService
func (boot *SimpleBootStarter) BootStartRpc() {
	// bootstart
	if boot.pakkuapp == nil {
		boot.builder.PakkuModules().EnableAppConfig().EnableAppService()
		boot.BootStart()
	}

	// 注册rpcservice
	service := boot.pakkuapp.PakkuModules().GetAppService()
	for _, v := range pakkuconf.RegisterRPCService(boot.pakkuapp) {
		if err := service.RegisteRPC(v); nil != err {
			logs.Panicln(err)
		}
	}

	// 读取配置&启动服务
	config := boot.pakkuapp.PakkuModules().GetAppConfig()
	service.StartRPC(ipakku.RPCServiceConfig{
		Debug:      false,
		ListenAddr: config.GetConfig(sysconstants.CONFIG_KEY_RPC_LISTEN_ADDRESS).ToString(sysconstants.DEFAULT_VAL_LISTEN_ADDRESS),
	})
}

// initialLogdir 初始化本地日志目录
func initialLogdir(logdir string) string {
	if len(logdir) == 0 {
		logdir = sysconstants.DEFAULT_VAL_LOG_DIR
	}
	if !filepath.IsAbs(logdir) {
		if path, err := filepath.Abs(logdir); nil != err {
			logs.Panicln(err)
		} else {
			logdir = path
		}
	}
	if !fileutil.IsExist(logdir) {
		if err := fileutil.MkdirAll(logdir); nil != err {
			logs.Panicln(err)
		}
	}
	return logdir
}

// getCertFile 获取https证书
func getCertFile() (certFile string, keyFile string) {
	if fileutil.IsFile(sysconstants.DEFAULT_VAL_CERT_KEY_FILE_PATH) {
		keyFile = sysconstants.DEFAULT_VAL_CERT_KEY_FILE_PATH
	}
	if fileutil.IsFile(sysconstants.DEFAULT_VAL_CERT_FILE_PATH) {
		certFile = sysconstants.DEFAULT_VAL_CERT_FILE_PATH
	}
	return
}
