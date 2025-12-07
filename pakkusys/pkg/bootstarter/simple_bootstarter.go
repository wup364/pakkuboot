// SPDX-License-Identifier: MIT
// Copyright (C) 2021 WuPeng <wup364@outlook.com>.

// 实例化一个pakku运行环境

package bootstarter

import (
	"pakkuboot/pakkusys"
	"pakkuboot/pakkusys/bootconfig"
	"path/filepath"
	"sync"

	"github.com/wup364/pakku"
	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/pkg/fileutil"
	"github.com/wup364/pakku/pkg/logs"
	"github.com/wup364/pakku/pkg/strutil"
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
		loggerWriter := bootconfig.RegisterLoggerWriter(logdir, logName, DEFAULT_VAL_MAX_LOG_FILES)

		logs.Info("Logger output file, path: ", logdir)
		boot.builder.PakkuConfigure().SetLoggerOutput(loggerWriter)
	}

	switch loglevel {
	case "none":
		boot.builder.PakkuConfigure().SetLoggerLevel(logs.LOG_LEVEL_NONE)
	case "error":
		boot.builder.PakkuConfigure().SetLoggerLevel(logs.LOG_LEVEL_ERROR)
	case "info":
		boot.builder.PakkuConfigure().SetLoggerLevel(logs.LOG_LEVEL_INFO)
	default:
		boot.builder.PakkuConfigure().SetLoggerLevel(logs.LOG_LEVEL_DEBUG)
	}
	return boot
}

// BootStart 加载&启动程序
func (boot *SimpleBootStarter) BootStart() ipakku.PakkuApplication {
	boot.locker.Lock()
	defer boot.locker.Unlock()

	if boot.pakkuapp != nil {
		return boot.pakkuapp
	}

	pakkuApp := boot.builder.Application()

	// 加载额外的自带模块
	bootconfig.EnablePakkuModules(boot.builder.PakkuModules())

	// 注册覆盖模块
	if overrides := bootconfig.RegisterOverride(pakkuApp); len(overrides) > 0 {
		for i := 0; i < len(overrides); i++ {
			if overrides[i].Instance != nil {
				ipakku.PakkuConf.RegisterPakkuModuleImplement(overrides[i].Instance, overrides[i].Interface, overrides[i].Implement)
			}
			ipakku.PakkuConf.SetPakkuModuleImplement(pakkuApp.Params(), overrides[i].Interface, overrides[i].Implement)
		}
	}

	// 加载自定义模块
	if modules := bootconfig.RegisterModules(pakkuApp); len(modules) > 0 {
		boot.builder.CustomModules().AddModules(modules...)
	}

	// 注册模块加载事件
	if events := bootconfig.RegisterModuleEvent(pakkuApp); len(events) > 0 {
		for i := 0; i < len(events); i++ {
			boot.builder.ModuleEvents().Listen(events[i].Module, events[i].Event, events[i].Handler)
		}
	}

	// 设置数据源到应用参数中
	boot.initDataSource(pakkuApp)

	boot.pakkuapp = boot.builder.BootStart()

	return boot.pakkuapp
}

// BootStartWeb 以web服务方式启动, https证书放在.conf目录下自动加载
// 自动装载模块: AppConfig AppService
func (boot *SimpleBootStarter) BootStartWeb(debug bool) {
	// bootstart
	if boot.pakkuapp == nil {
		boot.builder.PakkuModules().EnableAppConfig().EnableAppService()
		boot.BootStart()
	}

	// 注册controller
	service := boot.pakkuapp.PakkuModules().GetAppService()
	for _, v := range bootconfig.RegisterHttpController(boot.pakkuapp) {
		if err := service.AsController(v); nil != err {
			logs.Panic(err)
		}
	}

	// 注册http过滤器
	for _, v := range bootconfig.RegisterHttpRequestFilter(boot.pakkuapp) {
		if err := service.Filter(v.Path, v.Func); nil != err {
			logs.Panic(err)
		}
	}

	// 读取配置 & 启动服务
	certFile, keyFile := getCertFile()
	config := boot.pakkuapp.PakkuModules().GetAppConfig()
	service.StartHTTP(ipakku.HTTPServiceConfig{
		Debug:      debug,
		ListenAddr: config.GetConfig(CONFIG_KEY_HTTP_LISTEN_ADDRESS).ToString(DEFAULT_VAL_HTTP_LISTEN_ADDRESS),
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
	for _, v := range bootconfig.RegisterRPCService(boot.pakkuapp) {
		if err := service.RegisteRPC(v); nil != err {
			logs.Panic(err)
		}
	}

	// 读取配置&启动服务
	config := boot.pakkuapp.PakkuModules().GetAppConfig()
	service.StartRPC(ipakku.RPCServiceConfig{
		Debug:      false,
		ListenAddr: config.GetConfig(CONFIG_KEY_RPC_LISTEN_ADDRESS).ToString(DEFAULT_VAL_RPC_LISTEN_ADDRESS),
	})
}

// initDataSource 初始化数据源
func (boot *SimpleBootStarter) initDataSource(pakkuApp ipakku.PakkuApplication) {
	boot.builder.ModuleEvents().Listen(ipakku.ModuleID.AppConfig, ipakku.ModuleEventOnLoaded, func(module any, app ipakku.Application) {
		if dbIns, err := bootconfig.RegisterDataSource(pakkuApp, module.(ipakku.AppConfig)); nil != err {
			logs.Panic(err)
		} else if dbIns != nil {
			logs.Debugf("DataSource: %s, Driver: %s", dbIns.DataSourceUrl, dbIns.DriverName)
			boot.builder.Application().Params().SetParam(pakkusys.PAKKU_PARAMS_KEY_DATASOUCE, *dbIns)
		}
	})
}

// initialLogdir 初始化本地日志目录
func initialLogdir(logdir string) string {
	if strutil.IsBlank(logdir) {
		logdir = DEFAULT_VAL_LOG_DIR
	}
	if !filepath.IsAbs(logdir) {
		if path, err := filepath.Abs(logdir); nil != err {
			logs.Panic(err)
		} else {
			logdir = path
		}
	}
	if !fileutil.IsExist(logdir) {
		if err := fileutil.MkdirAll(logdir); nil != err {
			logs.Panic(err)
		}
	}
	return logdir
}

// getCertFile 获取https证书
func getCertFile() (certFile string, keyFile string) {
	if fileutil.IsFile(DEFAULT_VAL_CERT_KEY_FILE_PATH) {
		keyFile = DEFAULT_VAL_CERT_KEY_FILE_PATH
	}
	if fileutil.IsFile(DEFAULT_VAL_CERT_FILE_PATH) {
		certFile = DEFAULT_VAL_CERT_FILE_PATH
	}
	return
}
