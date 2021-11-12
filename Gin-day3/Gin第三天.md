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
   //8<<20 即 8*2^20
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



