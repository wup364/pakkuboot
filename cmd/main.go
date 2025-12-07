// SPDX-License-Identifier: MIT
// Copyright (C) 2021 WuPeng <wup364@outlook.com>.

// 主函数 - 加载并初始化服务

package main

import (
	"flag"
	"pakkuboot/pakkusys/pkg/bootstarter"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	name := flag.String("name", "pakkuapp", "App id, default pakkuapp")
	logger := flag.String("logger", "console", "logger: console, file or unset, default console")
	loglevel := flag.String("loglevel", "debug", "loglevel: debug, info, error or none, default debug")
	logdir := flag.String("logdir", "./logs", "default ./logs/{application name}.log")
	flag.Parse()

	// 启动应用&启用web服务
	appBoot := bootstarter.NewSimpleBootStarter(*name).SetLogger(*logger, *logdir, *loglevel)

	// 启动&&web服务
	appBoot.BootStartWeb(*loglevel == "debug")
}
