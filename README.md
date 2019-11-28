# gin-api       
基于Gin框架搭建，专门为app接口开发而设计的一个API统一网关。
  
Gin是一个用Go（Golang）编写的Web框架。
它具有类似马提尼的API，具有更好的性能，由于httprouter，速度提高了40倍。
如果你需要表现和良好的生产力，你会爱上Gin。

## 特性

- #### 快速
基于 Radix 树的路由，小内存占用。没有反射。可预测的 API 性能。

- #### 支持中间件
传入的 HTTP 请求可以由一系列中间件和最终操作来处理。 例如：Logger，Authorization，gzip，zaplog等，最终操作 DB。

- #### Crash 处理
Gin 可以 catch 一个发生在 HTTP 请求中的 panic 并 recover 它。这样，你的服务器将始终可用。例如，你可以向 Sentry 报告这个 panic！

- #### JSON 验证
Gin 可以解析并验证请求的 JSON，例如检查所需值的存在。

- #### 路由组
更好地组织路由。是否需要授权，不同的 API 版本…… 此外，这些组可以无限制地嵌套而不会降低性能。

- #### 错误管理
Gin 提供了一种方便的方法来收集 HTTP 请求期间发生的所有错误。最终，中间件可以将它们写入日志文件，数据库并通过网络发送。

- #### 内置渲染
Gin 为 JSON，XML 和 HTML 渲染提供了易于使用的 API。

- #### 可扩展性
新建一个中间件非常简单，去查看[示例代码](https://gin-gonic.com/zh-cn/docs/examples/using-middleware/)吧。

- #### 版本控制
轻松管理各个版本API，不用处理兼容问题，快速迭代

## 部署

### 要求
- Go 1.11 及以上版本

### 1. 进到项目根目录
```bash
$ cd gin-api
```

### 2. 执行bin目录下的 develop.sh 部署工具脚本文件
```bash
$ ./bin/develop.sh sync
$ ./bin/develop.sh restart
```

> 提醒事项：
- 此脚本运行方式只适合开发阶段

## 资料
- Gin官网：https://gin-gonic.com
- Github-Gin：https://github.com/gin-gonic
