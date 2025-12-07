// SPDX-License-Identifier: MIT
// Copyright (C) 2024 WuPeng <wup364@outlook.com>.

package bootstarter

const (
	// 配置KEY - http监听地址
	CONFIG_KEY_HTTP_LISTEN_ADDRESS = "pakku.listen.http.address"

	// 配置KEY - rpc监听地址
	CONFIG_KEY_RPC_LISTEN_ADDRESS = "pakku.listen.rpc.address"

	// DEFAULT_VAL_HTTP_LISTEN_ADDRESS 默认值 - http监听地址
	DEFAULT_VAL_HTTP_LISTEN_ADDRESS = "127.0.0.1:8080"

	// DEFAULT_VAL_RPC_LISTEN_ADDRESS 默认值 - rpc监听地址
	DEFAULT_VAL_RPC_LISTEN_ADDRESS = "127.0.0.1:8080"

	// DEFAULT_VAL_LOG_DIR 默认值 - 日志目录
	DEFAULT_VAL_LOG_DIR = "./logs"

	// DEFAULT_VAL_MAX_LOG_FILES 最大日志文件个数
	DEFAULT_VAL_MAX_LOG_FILES = 15

	// DEFAULT_VAL_CERT_KEY_FILE_PATH 默认值 - https证书信息
	DEFAULT_VAL_CERT_KEY_FILE_PATH = ".conf/key.pem"

	// DEFAULT_VAL_CERT_FILE_PATH 默认值 - https证书信息
	DEFAULT_VAL_CERT_FILE_PATH = ".conf/cert.pem"
)
