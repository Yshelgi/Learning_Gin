# Gin第三天

## 1.表单参数

表单传输为post请求，http常见的传输格式为四种：

-   application/json
-   application/x-www-form-urlencoded
-   application/xml
-   multipart/form-data

表单参数可以通过PostForm()方法获取，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数

首先我们写一个简单的提交表单，内容就是基本的用户名和密码

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>demo1</title>
</head>
<body>
<form action="http://localhost:8080/form" method="post" action="application/x-www-form-urlencoded">
    用户名：<input type="text" name="username" placeholder="请输入你的用户名">  <br>
    密&nbsp;&nbsp;&nbsp;码：<input type="password" name="userpassword" placeholder="请输入你的密码">  <br>
    <input type="submit" value="提交">
</form>
</body>
</html>
```

这里我们表单的action设定为等会post路径

```go
package main

import (
   "fmt"
   "github.com/gin-gonic/gin"
   "net/http"
)

//表单参数

func main(){
   r:=gin.Default()
   r.POST("/form", func(c *gin.Context) {
      types:=c.DefaultPostForm("type","post")
      // 键名和html页面属性名对应
      username:=c.PostForm("username")
      password:=c.PostForm("userpassword")
      c.String(http.StatusOK,fmt.Sprintf("username:%s,password:%s,types:%s",username,password,types))
   })
   r.Run()
}
```

![image-20211112192246221](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211112192246221.png)

![image-20211112192329990](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211112192329990.png)

## 2. 文件上传

### 2.1 单个文件上传

multipart/form-data格式用于文件上传，gin文件上传与原生的net/http方法类似，不同在于gin把原生的request封装到c.Request中。同样的，先写一个上传单个文件的HTML页面

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>demo2</title>
</head>
<body>
<form action="http://localhost:8080/upload" method="post" enctype="multipart/form-data">
  上传文件:<input type="file" name="file" >
  <input type="submit" value="提交">
</form>
</body>
</html>
```

然后用gin实现接收表单文件并保存

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

//单个文件上传
func main(){
   r:=gin.Default()
   //8<<20 即 8*2^20=8M
   r.MaxMultipartMemory=8<<20
   r.POST("/upload",func(c *gin.Context){
      file,err:=c.FormFile("file")
      if err!=nil{
         c.String(500,"上传文件出错")
      }
      c.SaveUploadedFile(file,file.Filename)
      c.String(http.StatusOK,file.Filename+"上传成功")
   })
   r.Run()
}
```

![image-20211112195132352](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211112195132352.png)

![image-20211112195146210](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211112195146210.png)

![image-20211112195231940](https://gitee.com/shixiaojiejiela_admin/pics/raw/master///image-20211112195231940.png)



如果我想限制只能上传png，不上传其他文件格式呢？我们可以自己实现限制文件上传类型的函数



从headers中选择文件类型，然后判断是否匹配。修改后的代码如下

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

//单个文件上传

func main() {
   r := gin.Default()
   //8<<20 即 8*2^20=8M
   r.MaxMultipartMemory = 8 << 20
   r.POST("/upload", func(c *gin.Context) {
      _,headers, err := c.Request.FormFile("file")
      if err != nil {
         c.String(500, "上传文件出错")
      }
      if headers.Header.Get("Content-Type")!="image/png"{
         c.String(500,"只能上传png文件")
         return
      }
      c.SaveUploadedFile(headers, "./imgs/"+headers.Filename)
      c.String(http.StatusOK, headers.Filename+"上传成功")
   })
   r.Run()
}
```



![image-20211112222837062](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211112222837062.png)

![image-20211112223332946](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211112223332946.png)

![image-20211112225231915](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211112225231915.png)

![image-20211112225247437](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211112225247437.png)

![image-20211112225857623](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211112225857623.png)



### 2.2 多文件上传

多文件上传就是一次可以上传多个文件，这也便于上传文件的人。（有的网站一次只能上传一个文件，属实太难受了）

上传的页面不用特别改动，只需要在input上传表单添加一个multiple

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>demo2</title>
</head>
<body>
<form action="http://localhost:8080/upload" method="post" enctype="multipart/form-data">
  上传文件:<input type="file" name="files" multiple>
  <input type="submit" value="提交">
</form>
</body>
</html>
```

主要是后端的操作，不过还好gin里面有多段表单模块(封装了"mime/multipart"，都是经过了解析之后的)

![image-20211112233525107](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211112233525107.png)

```go
package main

import (
   "fmt"
   "github.com/gin-gonic/gin"
   "net/http"
)

//多文件上传

func main() {
   r := gin.Default()
   r.MaxMultipartMemory = 8 << 20
   r.POST("/upload", func(c *gin.Context) {
      form, err := c.MultipartForm()
      if err != nil {
         c.String(http.StatusBadRequest, fmt.Sprintf("get err %s"), err.Error())
      }
      files := form.File["files"]
      for _,file :=range files{
         if err:=c.SaveUploadedFile(file,"./res/"+file.Filename);err!=nil{
            c.String(http.StatusBadRequest,fmt.Sprintf("upload err %s",err.Error()))
            return
         }
      }
      c.String(http.StatusOK,fmt.Sprintf("upload %d files",len(files)))
   })

   r.Run()
}
```

![image-20211112233121789](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211112233121789.png)

![image-20211112233929076](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211112233929076.png)

![image-20211112233958356](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211112233958356.png)





## 3. 404页面

往往我们开发的时候，是不知道未来用户会输入什么样的URL，所以我们还需要设置404页面，也就是找不到的页面

```go
package main

import (
   "fmt"
   "github.com/gin-gonic/gin"
   "net/http"
)

//设置404 NOT FOUND

func main(){
   r:=gin.Default()
   r.GET("/user",func(c *gin.Context){
      name:=c.DefaultQuery("name","shelgi")
      c.String(http.StatusOK,fmt.Sprintf("hello %s",name))
   })
   // 当访问到不知名路由
   r.NoRoute(func(c *gin.Context){
      c.String(http.StatusNotFound,"404 NOT FOUND")
   })
   r.Run()
}
```

先去我们已经设定好的路由界面

![image-20211112234701293](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211112234701293.png)

再试试未设置路由的界面

![image-20211112234730465](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211112234730465.png)

去看看NoRoute()

![image-20211112235035481](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211112235035481.png)

其实是引擎的一种方法，传入的是一个处理器函数，默认返回404

在最上面已经将这些处理器函数定义好了

![image-20211112235625998](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211112235625998.png)

再来看看引擎，盲猜肯定是包含很多方法和属性的结构体

![image-20211113000451026](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211113000451026.png)

果然，猜的没错。引擎其实就是框架的一个实例，包含了多路复用路由、中间件和设置，可以使用New()或者Default()创建一个引擎实例。这也就是为什么每次开头都是gin.Default()的原因所在了，**换言之我们也可以用gin.New()作为替换**。





## 学习Gin的一些体会

学了三天，就我的体验来说Gin上手难度并不大，只是不论是社区还是网络上的相关资料都不太丰富。我在之前接触过go语言，看完了基本经典的书籍以及web相关的书籍，不过没有涉及到具体的框架，**为了真正上手Gin（之前学过一丢丢就耽误了），又重新温习了一遍go语言的基础知识，我认为这在学Gin的时候是非常有用的**。

学习Gin，鉴于它资源少，官方文档不详细的问题，我们必须要学会自己去阅读理解源码，这也是为什么我的博客除了把文档中的内容复现之外添加许多源码解释部分的原因，这也要求我们对go有一定的掌握程度，所以学习任何框架前基础语言知识还是得搭稳（特别是对我这种经常一天内很多语言混用的人，思路清晰才不会混淆）

我也会坚持把Gin专栏写下去，尽量把一些细节讲清楚，从实现到源码解析希望能帮助到后来的学习者。另外，我目前还看上了一款框架---[GoFrame](https://goframe.org/display/gf)，国内开发的企业级框架，而且官方的文档非常丰富，Gin之后我应该就会开始上手GoFrame，就好比flask vs Django？



