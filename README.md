# pakkuboot 快速搭建服务模板

    使用pakku工具包封装的web服务快速开发实例, demo默认实现了一个syahello的http服务. 可在此模板的基础上开发新模块.

## 快速开发

    1. 拉取本项目到本地, 根据需要调整文件夹和文件名
    2. 在biz下通过编写service、controller等实现业务逻辑(默认demo: syahello.go)
    3. 在 `bootconfig.go` 中注册对应的对象(默认demo: syahello.go)
    4. 运行或编译 'main.go`
    5. 浏览器访问 'http://127.0.0.1:8080/sayhello/v1/hello'

### 目录说明

|  目录 |  文件  |  描述  |
| ----- | ------ | ------ |
| / | main.go | 程序入口, 默认启动实现方法 |
| /biz | * | 业务代码位置 |
| /pakkusys/application | * | 实例化pakku的工具类, 启动工具 |
| /pakkusys/pakkuconf | * | 类配置目录, 存放一些配置性文件, 如: 使用redis、kafka实例替代系统默认的cache、event实现 |
| /pakkusys/pakkuconf | bootconfig.go | 在`RegisterModules`方法内注册新模块、`RegisterOverride`内设置覆盖的模块 |
| /.conf | key.pem,cert.pem | 默认https证书存放位置, 默认没有 |

### 配置清单(默认)

|  KEY  |  默认值  |  可选值  |  描述 |
| ------ | ------ | ---- | ---- |
| `listen.http.address` | 127.0.0.1:8080 | `*` | 对外HTTP服务地址(默认启用) |
| `listen.rpc.address` | 127.0.0.1:8080 | `*` | 对外RPC服务地址(默认不启用) |

    . 配置使用json格式存储, 格式示例: 
        `{
            "listen": {
                "http": {
                    "address": "127.0.0.1:5051"
                }
            }
        }` => listen.http.address

### 模块的编写

    通常围绕着`Module`、`Controller`编写业务, biz下包含一个`SayHello`的服务和controller.

## pakku 帕克概述

    pakku 的核心是提供一个对象加载的环境, 在加载对象的同时提供额外的服务, 如依赖注入等操作. 其核心轻量无且三方引用依赖. 使用内置的几个模块, 便可快速搭建一个简单的服务或程序.

### 内置的接口

    pakku 在加载实现了`ipakku.Module`、`ipakku.Controller`、`ipakku.Router`接口的对象以及RPC注册对象时, 会对内部的成员变量(包括私有)进行扫描, 对包含`@autowired:"接口名"`标签的字段进行自动赋值, 实现了模块间的解耦和依赖注入.

* `Module` 即模块, 通常为有一定功能单元的对象, 在`Controller`、`rpc服务`、`其他Module`中被注入和调用.
* `Router` 即http路由定义模块, 通过定义接口的url、地址列表实现对url的路由.
* `Controller` 即http服务定义模块, 通过定义接口的url、地址列表、过滤器等实现对url的路由.
* `RpcService` 即rpc服务定义模块, 默认使用自带的rpc服务进行注册.

### 内置的模块

    pakku 默认实现了AppConfig(配置模块)、AppCache(缓存模块)、AppEvent(事件模块)、AppService(NET服务模块)以满足一个基本的服务运行环境. 如需使用默认的接口定义但又需要使用其他方式实现, 可通过重新实现对应`ipakku/Ixxx`的接口后, 重新指定默认调用实例即可(ipakku.Override.RegisterInterfaceImpl).

|  名字 |  可重写接口类  |  描述  |
| ------ | ------ | ------ |
| AppConfig | `ipakku.IConfig` | 使用json格式存储的配置实现, 文件存放在启动目录下`./.conf/{appName}.json`中 |
| AppCache | `ipakku.ICache` | 使用map实现的本地内存缓存, 如需使用其他缓存机制, 如redis需要自己实现 |
| AppEvent | `ipakku.IEvent` | 默认没有实现此接口, 需要自己实现, 如: kafka等 |
| AppService | `-` | 默认实现了http服务和rpc服务, 不可重写 |
