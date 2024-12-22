// Copyright (C) 2024 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package sysconstants

const (
	// C_DB_TYPE_SQLITE3 数据库类型 - sqlite3
	C_DB_TYPE_SQLITE3 = "sqlite3"

	// C_MAX_LOG_FILES 最大日志文件个数
	C_MAX_LOG_FILES = 15
)

const (
	// 配置KEY - http监听地址
	CONFIG_KEY_HTTP_LISTEN_ADDRESS = "listen.http.address"

	// 配置KEY - rpc监听地址
	CONFIG_KEY_RPC_LISTEN_ADDRESS = "listen.rpc.address"

	// CONFIG_KEY_DATASOURCE_DRIVER 配置KEY - 数据源驱动(sqlite3、mysql)
	CONFIG_KEY_DATASOURCE_DRIVER = "datasource.driver"

	// CONFIG_KEY_DATASOURCE_URL 配置KEY - 数据源URL
	CONFIG_KEY_DATASOURCE_URL = "datasource.url"

	// DEFAULT_VAL_LISTEN_ADDRESS 默认值 - http|rpc监听地址
	DEFAULT_VAL_LISTEN_ADDRESS = "127.0.0.1:8080"

	// DEFAULT_VAL_LOG_DIR 默认值 - 日志目录
	DEFAULT_VAL_LOG_DIR = "./logs"

	// DEFAULT_VAL_CERT_KEY_FILE_PATH 默认值 - https证书信息
	DEFAULT_VAL_CERT_KEY_FILE_PATH = ".conf/key.pem"

	// DEFAULT_VAL_CERT_FILE_PATH 默认值 - https证书信息
	DEFAULT_VAL_CERT_FILE_PATH = ".conf/cert.pem"

	// DEFAULT_VAL_DATAS_DIR 默认值 - 数据目录
	DEFAULT_VAL_DATAS_DIR = ".datas"

	// DEFAULT_VAL_DATASOURCE_DRIVER 默认值 - 数据源驱动
	DEFAULT_VAL_DATASOURCE_DRIVER = C_DB_TYPE_SQLITE3

	// DEFAULT_VAL_DATASOURCE_URL 默认值 - 数据源URL
	DEFAULT_VAL_DATASOURCE_URL = DEFAULT_VAL_DATAS_DIR + "/%s.db?cache=shared"
)
