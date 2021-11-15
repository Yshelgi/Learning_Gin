# Gin第四天



## 1. 路由组

当我们有许多相同的URL，但是需要它们处理不同的请求，这个时候就可以利用路由组来分别设置。

例如把GET请求的路由放在一个分组里，把POST请求的分组放在另一个。

```go
package main

import (
   "fmt"
   "github.com/gin-gonic/gin"
   "net/http"
)

func main(){
   r:=gin.Default()
   v1:=r.Group("/get")
   {
      v1.GET("/login",login)
      v1.GET("/submit",submit)
   }

   v2:=r.Group("/post")
   {
      v2.POST("/login",login)
      v2.POST("/submit",submit)
   }
   r.Run()
}

func login(c *gin.Context){
   name:=c.DefaultQuery("name","shelgi")
   c.String(http.StatusOK,fmt.Sprintf("hello %s\n",name))
}

func submit(c *gin.Context){
   name:=c.DefaultQuery("name","shelgi")
   c.String(http.StatusOK,fmt.Sprintf("hello %s\n",name))
}
```

![image-20211115131248201](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211115131248201.png)

![image-20211115131331407](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211115131331407.png)

但如果我们用post请求去请求get操作对应的页面，就会返回404

![image-20211115131518151](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211115131518151.png)

来看看Group到底干了什么，它是RouterGroup的一个方法，然后内部将传入的相对路径转换为绝对路径，并且把处理函数结合起来（下图二），返回一个路由组

![image-20211115131654881](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211115131654881.png)

![image-20211115132116999](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211115132116999.png)

可以看到这里对传入的处理函数有规模限制，这里面abortIndex从源码可以看到

`const abortIndex int8=math.MaxInt8 / 2`

也就是总共的处理函数不能超过**63个**

关于Gin的路由原理，涉及到前缀树和基数树，**基数树就是前缀树压缩优化后的结果**；Gin采用httprouter进行路由匹配，每个http方法对应都会生成一颗基数树，下面是树节点的结构

![image-20211115134819500](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211115134819500.png)

里面包含了路径、索引、前节点、孩子节点、处理函数链……

感兴趣的可以去gin源码中的tree.go看看实现



## 2. 路由的拆分与注册

### 2.1 基本的路由注册

最基本的路由注册就是先写一个处理函数，然后再将这个处理函数与请求方法与路由路径绑定上，类似于下面的例子

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

func testhandler(c *gin.Context){
   c.JSON(http.StatusOK,gin.H{
      "test":"这是路由注册的测试",
   })
}

func main(){
   r:=gin.Default()
   r.GET("/test",testhandler)
   r.Run()
}
```

![image-20211115140443027](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211115140443027.png)



### 2.2 路由拆分为包

当我们需要处理注册的路由数量过多时，全部写在一个文件里既不方便阅读修改，也不能很好的简化代码结构，为此我们可以把这部分代码拎出来单独作为一个包

创建一个routers的包，然后将路由注册部分全部放在这个包里

```go
package routers

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

func testhandler(c *gin.Context){
   c.JSON(http.StatusOK,gin.H{
      "test":"这是路由注册的测试",
   })
}

func SetupRouter() *gin.Engine{
   r:=gin.Default()
   r.GET("/test",testhandler)
   return r
}
```

然后在主函数中引入已经实例化路由注册后的引擎

```go
package main

import (
   "Learn_Gin/Gin-day4/routers"
)

//从Routers包中引入路由

func main() {
   r := routers.SetupRouter()
   r.Run()
}
```

![image-20211115141820768](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211115141820768.png)



### 2.3 路由拆分为多个文件

我们现在是把所有的路由注册全部写在SetupRouter中，但是如果我们的路由更多，同样会引起之前一样的问题，所以我们还可以继续分，分为更细致的模块化

在routers中创建两个路由，传入引擎实例作为参数，在模块中就实现了各个路由的注册

```go
package routers

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

func test1handler(c *gin.Context){
   c.JSON(http.StatusOK,gin.H{
      "test":"这是路由1注册的测试",
   })
}

func Test1(e *gin.Engine){
   e.GET("/test1",test1handler)
}
```

```go
package routers

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

func test2handler(c *gin.Context){
   c.JSON(http.StatusOK,gin.H{
      "test":"这是路由2注册的测试",
   })
}


func Test2(e *gin.Engine){
   e.GET("/test2",test2handler)
}
```

在主函数中只用调用这两个函数就可以啦

```go
package main

import (
   "Learn_Gin/Gin-day4/routers"
   "github.com/gin-gonic/gin"
)

//多文件路由注册

func main(){
   r:=gin.Default()
   routers.Test1(r)
   routers.Test2(r)
   r.Run()
}
```

![image-20211115143810869](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211115143810869.png)

![image-20211115143824276](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211115143824276.png)

说到这里，来互相对比一下。拿我们熟悉的**Flask**作为对比，除了`@app.route("/")` 装饰器方式注册路由，还可以使用`app.add_url_rule(rule="xxx",view_func=func)`

这个和我们的`r.GET(relativePath="/",handles)`结构上其实是一样的，有的时候交叉对比学习，可以进行知识迁移，从而更简单的理解使用



### 2.4 路由拆分为多个APP

当我们的项目继续扩大，全部放在routers里面还是不能满足分模块的效果，我们就可以继续再分；每个不同的功能模块放在APP的不同模块下，然后在各自的模块中实现handler和Routers（路由注册），最终只需要在routers中将所有APP注册的路由全部整合在一起（一个切片中），然后初始化方法中用实例一个个遍历注册。在主函数中只需要直接初始化然后运行就可以了。

APP/demo/下

```go
package demo

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

func idHandler(c *gin.Context){
   c.JSON(http.StatusOK,gin.H{
      "id_test":"这是id的测试",
   })
}

func commentHandler(c *gin.Context){
   c.JSON(http.StatusOK,gin.H{
      "comment_test":"这是comment的测试",
   })
}
```

```go
package demo

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
   e.GET("/id", idHandler)
   e.GET("/comment", commentHandler)
}
```

routers中合并注册整合，然后同意初始化

```go
package routers

import "github.com/gin-gonic/gin"

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
   options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
   r := gin.New()
   for _, opt := range options {
      opt(r)
   }
   return r
}
```

主函数中调用整合初始化方法，直接启动

```go
package main

import (
   "Learn_Gin/Gin-day4/APP/demo"
   "Learn_Gin/Gin-day4/routers"
)

func main(){
   routers.Include(demo.Routers)
   r:=routers.Init()
   r.Run()
}
```

![image-20211115160636825](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211115160636825.png)



### 2.5 总结

这部分主要就是在处理一件事，如何解耦，使各个模块之间更细分并且在项目方便使用。不论是从项目的目录结构还是从代码功能都能让人比较清晰直观的理解，这样的项目不仅利于开发，更利于后期的迭代更新、运营维护。

