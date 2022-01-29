# Gin第七天

Gin的中间件，或者说web框架中的中间件是提供系统软件和应用软件之间连接的软件。这些中间件可以是第三方的，也可以是自定义的，它们可以为web程序添加功能，也可以多处复用。

在使用Gin创建路由的时候，我们经常会用到`gin.Default()`，其实这个之前就提到过，它默认就使用了Logger()和Recovery()这两个中间件。

![image-20220129221859360](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20220129221859360.png)

如果不想使用这两个可以直接使用`gin.New()`，其实一般情况下这两种差别不大。今天就主要讲讲Gin中关于中间件的使用问题，其中中间件又可以分为**全局中间件**和**局部中间件**。



## 1. 全局中间件

全局中间件就是整个项目中都可以随处调用，所有请求都要经过这个中间件。

我们可以写一个计时中间件的例子

```go
package main

import (
   "fmt"
   "github.com/gin-gonic/gin"
   "time"
)

func MiddleWare() gin.HandlerFunc{
   return func(c *gin.Context) {
      t:=time.Now()
      fmt.Println("全局中间件开始")
      c.Set("request","中间件")
      c.Next()
      status:=c.Writer.Status()
      fmt.Println("中间件执行完毕",status)
      t2:=time.Since(t)
      fmt.Println("耗时:",t2)
   }
}


func main(){
   r:=gin.Default()
   r.Use(MiddleWare())

   r.GET("test", func(c *gin.Context) {
      req,_:=c.Get("request")
      fmt.Println("request:",req)
      c.JSON(200,gin.H{"request":req})
   })

   r.Run()
}
```

可以看到，全局中间件通过`r.Use()`方法设置，所以最简单的例子

```go
r:=gin.Default()

r:=gin.New()
r.Use(Logger(),Recovery())
```

上面两种写法实际上是等价的，这从源码中也可以看出来。

设置完中间件，在中间件中如果要传递数据那么就需要使用到`gin.Context`中的set、get方法，将数据存放到Context中进行传递。

![image-20220129223409154](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20220129223409154.png)



## 2. 局部中间件

局部中间件就只有相对应的作用域中使用，只有相应路由请求才会使用到这个中间件。

同样的，也是写一个例子

```go
package main

import (
   "fmt"
   "github.com/gin-gonic/gin"
   "time"
)

func MiddleWare1() gin.HandlerFunc{
   return func(c *gin.Context) {
      t:=time.Now()
      fmt.Println(c.Request.URL.Path+"局部中间件开始")
      c.Set("request","中间件")
      c.Next()
      status:=c.Writer.Status()
      fmt.Println("局部中间件执行完毕",status)
      t2:=time.Since(t)
      fmt.Println("耗时:",t2)
   }
}


func main(){
   r:=gin.Default()

   r.GET("test",MiddleWare1(), func(c *gin.Context) {
         req,_:=c.Get("request")
         fmt.Println("request:",req)
         c.JSON(200,gin.H{"request":req})
      })

   r.Run()
}
```

我自定义了一个计时中间件，但是这次我不想全局设置，而是只用于个别路由，那么设置方式就有所改变。

![image-20220129223738275](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20220129223738275.png)

Gin中已经将所有的请求参数设为可变参数，也就是说除了路由，后面可以有很多处理器函数，中间件其实也就是一种处理器函数，所以我们可以像上面的例子一样，将局部中间件作为参数设置到对应的路由请求中，当然也可以像下面这种

`r.GET(...).Use(MiddleWare1())`

![image-20220129224423360](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20220129224423360.png)

源码中同样为路由组写了接口，所有路由请求都可以使用Use()方法，本质上其实就是使用append将路由组中的处理器函数进行添加然后返回新的路由组。

## 3. 案例+写入日志文件

有了上面的了解，我们就可以模拟一下中间件的使用，同时会涉及一些日志文件写入的问题。

```go
package main

import (
   "fmt"
   "github.com/gin-gonic/gin"
   "io"
   "os"
   "time"
)

func myMiddle(c *gin.Context){
   s:=time.Now()
   c.Next()
   e:=time.Since(s)
   fmt.Println("耗时:",e)
}


func handler1(c *gin.Context){
   time.Sleep(1*time.Second)
   c.JSON(200,gin.H{"handler1":"完成"})
}


func handler2(c *gin.Context){
   time.Sleep(2*time.Second)
   c.JSON(200,gin.H{"handler2":"完成"})
}

func main(){
   f,_:=os.Create("demo.log")
   gin.DefaultWriter=io.MultiWriter(f,os.Stdout)


   r:=gin.Default()
   r.Use(myMiddle)
   Group:=r.Group("/test")
   {
      Group.GET("/test1",handler1)
      Group.GET("/test2",handler2)
   }
   r.Run()
}
```

这里设置了两个handler来模拟一些请求事件，然后依然是一个计时中间件，但是在主函数中首先创建了一个log文件来记录日志，然后将输出流和文件流同时作为Gin的输出方式，全局注册中间件之后启动服务，访问一下不同路由。

![image-20220129230022875](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20220129230022875.png)

![image-20220129230043459](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20220129230043459.png)

这样就可以将日志记录在文件中。



## 结束

今天主要讲了在Gin中使用中间件，其实最近在学习fastapi中最大的兴奋之处还不是惊讶于python的异步框架性能能达到这种程度，而是服务启动后在/docs中可以直接查看api接口并且可以简单测试这个功能，对于前后端分离调用接口的时候感觉很舒服，还有一点关于Vue组件的东西没弄完，过年继续。
