# Gin框架入门---第一天

## 1.介绍

**Gin**是一个用Go语言编写的web框架，类似于martini但是拥有更好性能的API框架。

### 1.1 特性

+ 快速：基于Radix树的路由，小内存占用，没有反射，可预测的API性能
+ 支持中间件：传入的HTTP请求可以由一系列中间件和最终的操作来处理。例如：Logger、GZIP……
+ Crash 处理：Gin 可以 catch 一个发生在 HTTP 请求中的 panic 并 recover 它。这样，你的服务器将始终可用。例如，你可以向 Sentry 报告这个 panic！
+ JSON 验证：Gin 可以解析并验证请求的 JSON，例如检查所需值的存在。
+ 路由组：更好地组织路由。是否需要授权，不同的 API 版本…… 此外，这些组可以无限制地嵌套而不会降低性能。
+  错误管理：Gin 提供了一种方便的方法来收集 HTTP 请求期间发生的所有错误。最终，中间件可以将它们写入日志文件，数据库并通过网络发送。
+  内置渲染：Gin 为 JSON，XML 和 HTML 渲染提供了易于使用的 API。
+ 可扩展性：新建一个中间件非常简单



### 1.2 安装

安装Gin框架可以去按照[Go语言gin框架的安装_shelgi的博客-CSDN博客_gin框架下载](https://blog.csdn.net/shelgi/article/details/103940413)

引入Gin也非常简单：`import "github.com/gin-gonic/gin"`

当我们需要使用http.StatusOK等响应常量时，还需要额外引入net/http

`import net/http`



main.go

```go
package main

import "github.com/gin-gonic/gin"

func main() {
   r := gin.Default()
   r.GET("/ping", func(c *gin.Context) {
      c.JSON(200, gin.H{
         "message": "Hello Gin!",
      })
   })
   r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
```

![image-20211106231210733](/Users/ysj/Library/Application Support/typora-user-images/image-20211106231210733.png)



## 2. Gin的热启动

我查了一圈，Gin框架本身好像没有支持热启动的功能。那既然本身不支持，那我们就去寻求外界帮助，查找后发现大多数解决办法就是通过air，但是需要安装air，并且还需要在项目里去配置air.conf。我是一个比较懒的人，永远都去想选择最简单高效的方法实现需求，如果你也是这样的人，可以看看下面的办法

```go
github地址:https://github.com/codegangsta/gin
```

**gin**是用于实时重新加载**Go Web**应用程序的简单命令行实用程序。只需gin在应用程序目录中运行，网络应用程序将 gin作为代理提供。gin检测到更改后，将自动重新编译您的代码。您的应用在下次收到HTTP请求时将重新启动。不过如果编译出错或者发生错误时，它还是会“罢工”



**安装**

`go get github.com/codegangsta/gin`

安装结束后，我们测试一下是否安装成功

`gin -h`

![image-20211106235105429](/Users/ysj/Library/Application Support/typora-user-images/image-20211106235105429.png)

然后尝试一下热部署刚才的代码，启动服务

`gin run main.go`

![](/Users/ysj/Library/Application Support/typora-user-images/image-20211106235350460.png)

发现gin的默认端口是3000

然后把message改为Hello Shelgi，刷新一下网页

![image-20211106235544298](/Users/ysj/Library/Application Support/typora-user-images/image-20211106235544298.png)

这就是第一天的内容了，简单容易上手，后面再慢慢把框架使用讲清楚。