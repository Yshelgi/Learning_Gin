# Gin第五天

通过后端需要接收前端页面传来的参数，然后进行解析。传来的数据类型可能是json或者其他数据类型，分别看看几种不同的处理方式。

## 1.Json数据解析和绑定

创建一个处理Json格式的路由，其中接收数据的格式已经提前定义为一种结构体。

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

// 定义接收数据的结构体
type Login struct {
   // binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
   User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
   Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
   // 1.创建路由
   // 默认使用了2个中间件Logger(), Recovery()
   r := gin.Default()
   // JSON绑定
   r.POST("loginJSON", func(c *gin.Context) {
      // 声明接收的变量
      var json Login
      // 将request的body中的数据，自动按照json格式解析到结构体
      if err := c.ShouldBindJSON(&json); err != nil {
         // 返回错误信息
         // gin.H封装了生成json数据的工具
         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
         return
      }
      // 判断用户名密码是否正确
      if json.User != "root" || json.Pssword != "admin" {
         c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
         return
      }
      c.JSON(http.StatusOK, gin.H{"status": "200"})
   })
   r.Run()
}
```

![image-20211126165743941](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211126165743941.png)

![image-20211126165804133](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211126165804133.png)

![image-20211126170133015](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211126170133015.png)

其中，`c.json()`就是用json格式返回响应

![image-20211126170303919](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211126170303919.png)

源码也很简单，将传入的数据转为json格式输出到页面

[关于ShouldBindJSON](https://blog.csdn.net/heart66_A/article/details/100796964)

这里面讲了ShouldBind以及Bind等的区别，大家可以看看。



## 2. 表单数据解析和绑定

上面的是传入json数据，但是更多场景下，我们都是前端上传入表单数据给后端服务，所以更多还要处理表单数据。

先写一个简单的前端表单html

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
<form action="http://localhost:8000/loginForm" method="post" enctype="application/x-www-form-urlencoded">
    用户名<input type="text" name="username"><br>
    密码<input type="password" name="password">
    <input type="submit" value="提交">
</form>
</body>
</html>
```

服务端代码

```go
package main

import (
   "net/http"

   "github.com/gin-gonic/gin"
)

// 定义接收数据的结构体
type Login1 struct {
   // binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
   User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
   Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
   // 1.创建路由
   // 默认使用了2个中间件Logger(), Recovery()
   r := gin.Default()
   // JSON绑定
   r.POST("/loginForm", func(c *gin.Context) {
      // 声明接收的变量
      var form Login1
      // Bind()默认解析并绑定form格式
      // 根据请求头中content-type自动推断
      if err := c.Bind(&form); err != nil {
         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
         return
      }
      // 判断用户名密码是否正确
      if form.User != "root" || form.Pssword != "admin" {
         c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
         return
      }
      c.JSON(http.StatusOK, gin.H{"status": "200"})
   })
   r.Run()
}
```

![image-20211126172447845](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211126172447845.png)

![image-20211126172508109](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211126172508109.png)



## 3. URI数据解析和绑定

还有一种情况，我们的参数全部就在我们的URI中，这种情况我们可以解析路由。联系到第二天我们讲了路由api中的通配符，其实这里就要用到。

```go
package main

import (
   "net/http"

   "github.com/gin-gonic/gin"
)

// 定义接收数据的结构体
type Login2 struct {
   // binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
   User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
   Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
   // 1.创建路由
   // 默认使用了2个中间件Logger(), Recovery()
   r := gin.Default()
   // JSON绑定
   r.GET("/:user/:password", func(c *gin.Context) {
      // 声明接收的变量
      var login Login2
      // Bind()默认解析并绑定form格式
      // 根据请求头中content-type自动推断
      if err := c.ShouldBindUri(&login); err != nil {
         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
         return
      }
      // 判断用户名密码是否正确
      if login.User != "root" || login.Pssword != "admin" {
         c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
         return
      }
      c.JSON(http.StatusOK, gin.H{"status": "200"})
   })
   r.Run(":8000")
}
```

![image-20211126173331643](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211126173331643.png)

我们在GET方法中就用到了通配符去匹配对应的参数，然后用Gin自带的解析器去解析和绑定



## 总结

其实框架已经帮我们把大部分解析绑定的任务实现了，我们只需要事先定义好接收数据的结构体，以及结构体中对应不同格式的字段名称，然后就可以实现解析和绑定。

