# pakkuboot 快速搭建服务模板

    使用pakku工具包封装的web服务快速开发实例, demo默认实现了一个UserManagement的http服务. 可供参考.

## 快速开发

    1. 拉取本项目到本地, 根据需要调整文件夹和文件名
    2. 在business下通过编写service、controller等实现业务逻辑
    3. 在 `bootconfig.go` 中注册对应的对象
    4. 运行或编译 'main.go`
    5. 浏览器访问 'http://127.0.0.1:8080/user/v1'

### 目录说明

|  目录 |  文件  |  描述  |
| ----- | ------ | ------ |
| / | main.go | 程序入口, 默认启动实现方法 |
| /business | * | 业务代码位置 |
| /pakkusys/bootstarter | * | 实例化pakku的工具类, 启动工具 |
| /pakkusys/pakkuconf | * | 类配置目录, 存放一些配置性文件, 如: 使用redis、kafka实例替代系统默认的cache、event实现 |
| /pakkusys/pakkuconf | bootconfig.go | 在`RegisterModules`方法内注册新模块、`RegisterOverride`内设置覆盖的模块 |
| /.conf | key.pem,cert.pem | 默认https证书存放位置, 默认没有 |

### 配置清单(默认)

|  KEY  |  默认值  |  可选值  |  描述 |
| ------ | ------ | ---- | ---- |
| `listen.http.address` | :8080 | `*` | 对外HTTP服务地址(默认启用) |
| `listen.rpc.address` | :8080 | `*` | 对外RPC服务地址(默认不启用) |

    . 配置使用json格式存储, 格式示例: 
        `{
            "listen": {
                "http": {
                    "address": "127.0.0.1:5051"
                }
            }
        }` => listen.http.address

### 模块的编写

    通常围绕着`Module`、`Controller`编写业务, business下包含一个`UserManagement`的服务和controller.

## pakku 帕克概述

    pakku 的核心是提供一个对象加载的环境, 对类型为`ipakku.Module`的对象进行加载, 通过对各个`Module`加载节点事件的监听, 提供额外的辅助功能. 如依赖注入、自动完成配置值等操作. 其核心轻量无且三方引用依赖. 使用内置的几个模块, 便可快速搭建一个简单的服务或程序.
    

### 内置的模块

    pakku 默认实现了AppConfig(配置模块)、AppCache(缓存模块)、AppEvent(事件模块)、AppService(NET服务模块)以满足一个基本的服务运行环境. 如需使用默认的接口定义但又需要使用其他方式实现, 可通过重新实现对应`ipakku/Ixxx`的接口后, 重新指定默认调用实例即可(ipakku.Override.RegisterPakkuModuleImplement).

|  名字 |  可重写接口类  |  描述  |
| ------ | ------ | ------ |
| AppConfig | `ipakku.IConfig` | 使用json格式存储的配置实现, 文件存放在启动目录下`./.conf/{appName}.json`中 |
| AppCache | `ipakku.ICache` | 使用map实现的本地内存缓存, 如需使用其他缓存机制, 如redis需要自己实现 |
| AppEvent | `ipakku.IEvent` | 默认没有实现此接口, 需要自己实现, 如: kafka等 |
| AppService | `-` | 默认实现了http服务和rpc服务, 不可重写, 但可选是否启用该模块 |


### 特殊标签(tag)

    通过标注在struct的特殊标签值, 来实现一些辅助功能. 

|  TAG |  所属模块  |  作用域  |  格式  |  描述  |
| ------ | ------ | ------ | ------ | ------ |
| `@autowired` | Loader(加载器) | struct成员字段 | `@autowired:"模块名"` | 通过指定该标签, 可实现依赖对象自动注入 |
| `@autoConfig` | AppConfig(配置模块) | struct成员字段 | `@autoConfig:"配置路径"` |  标注当前字段是个配置struct, 可选'配置路径'参数  |
| `@value` | AppConfig(配置模块) | struct成员字段 | `@value:"配置路径"` | 通过'配置路径'查找并自动赋值对应字段, 可选'配置路径'参数 |
