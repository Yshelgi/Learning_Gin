# Gin第六天

## 1.数据响应

我们后端可以实现针对不同路由对应不同数据格式的响应，常用的是json格式，当然还有许多我们可以返回的数据格式，例如结构体、XML……

```go
package main

import(
   "github.com/gin-gonic/gin"
   "github.com/gin-gonic/gin/testdata/protoexample"
)


func main(){
   r:=gin.Default()

   // JSON格式
   r.GET("/JSON", func(c *gin.Context) {
      c.JSON(200,gin.H{"message":"someJSON","status":200})
   })

   // 结构体格式
   r.GET("/Struct", func(c *gin.Context) {
      var msg struct{
         Name string
         Message string
         Number int
      }

      msg.Name="test"
      msg.Message="message"
      msg.Number=123
      c.JSON(200,msg)
   })

   // XML
   r.GET("/XML", func(c *gin.Context) {
      c.XML(200,gin.H{"message":"abc"})

   })

   // YAML
   r.GET("/YAML", func(c *gin.Context) {
      c.YAML(200,gin.H{"name":"test"})
   })

   // protobuf
   r.GET("/Protobuf", func(c *gin.Context) {
      resp:=[]int64{int64(10),int64(2)}
      label:="label"
      data:=&protoexample.Test{
         Label: &label,
         Reps: resp,
      }
      c.ProtoBuf(200,data)
   })


   r.Run(":8888")
}
```

Gin中已经将常用的数据格式类型封装好了，只需要按要求将想要返回的数据格式内容写好



## 2. 模板渲染

这个是几乎所有web框架离不开的一个功能，也是现在很少用到的功能。为什么这么说呢，我学习了很多种web开发的框架，几乎都涉及到了html的模板渲染(模版引擎)，像python里面的jinja2、springboot中的thymeleaf以及gin的LoadHTMLFiles()。但是实际开发中大多选择前后端分离，很少在用模板渲染。所以这里只是简单示范

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

func main(){
   r:=gin.Default()
   r.LoadHTMLGlob("static/*")
   r.GET("/index", func(c *gin.Context) {
      c.HTML(http.StatusOK,"index.html",gin.H{"title":"我是测试","ce":"123456"})
   })
   r.Run()
}
```

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.title}}</title>
</head>
<body>
asdfjdsjklfsdlkfj{{.ce}}
</body>
</html>
```

![image-20220127233229313](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20220127233229313.png)

{{}}中就可以将路由参数作为变量传入前端页面进行渲染



## 3.重定向

重定向应该也是web开发中比较常用到的一个功能，前端经常可以写a标签点击跳转，后端如果想要将访问的路由修改到新的路由，也需要使用重定向。假如我们想将自己的web应用首页设置为百度，就可以像下面这样重定向到百度首页

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

func main(){
   r:=gin.Default()
   r.GET("/index", func(c *gin.Context) {
      c.Redirect(http.StatusMovedPermanently,"https://www.baidu.com")
   })

   r.Run()
}
```

![image-20220127234230708](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20220127234230708.png)

返回的是301，也就是永久重定向。这里稍微提一下，响应码302对应临时重定向，二者的区别在于永久重定向代表浏览器缓存了这个重定向的关联，下次请求相同路由的时候会直接发送重定向后的路由请求。这也就是为什么很多时候明明修改了重定向的路由，但是浏览器打开显示的还是之前的重定向页面。解决这个问题很简单，可以更换一种浏览器打开，或者清除浏览器缓存。



## 4. 同步与异步

众所周知，Go受欢迎一个主要原因就是与生俱来的并发优势，在很多时候go对于异步处理会比其他语言简单很多。对于I/O密集型的任务，使用多线程可以大大提高CPU使用率。在go语言中，我们只需要一个go关键字，就可以将函数变为异步执行，但是注意在使用新的goroutine时我们不应该使用原始的上下文，而是应该使用它的副本。同时要注意多线程之间的并发安全问题。

```go
package main

import (
   "github.com/gin-gonic/gin"
   "log"
   "time"
)

func main(){
   r:=gin.Default()
   r.GET("/async", func(c *gin.Context) {
      copyContext:=c.Copy()
      go func(){
         time.Sleep(3*time.Second)
         log.Println("异步执行:"+copyContext.Request.URL.Path)
      }()
   })

   r.GET("/sync", func(c *gin.Context) {
      time.Sleep(3*time.Second)
      log.Println("同步执行:"+c.Request.URL.Path)
   })

   r.Run()
}
```

![image-20220127235938768](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20220127235938768.png)

这样就对比了同步和异步执行的响应时间，可以看到差距还是比较明显的。



## 结尾

最近除了Gin，也在看python的fastapi，同样是异步支持框架，感觉写法上和flask很相似，所以学习成本不会太高，主要就是异步编码时async/await的使用，而且目前貌似fastapi社区中关于异步支持的插件并不是特别多，而且官方文档有的地方中文版都没有翻译完，很多时候看原版文档有些词不翻译还是不好理解。最后慢慢也要把Vue2收尾，再慢慢转向Vue3，争取假期还是能学完一点东西的。