# Gin第二天

## 1.RESTful API

简单来说，REST的含义就是客⼾端与Web服务器之间进⾏交互的时候，使⽤HTTP协议中的4个请求⽅法代表不同的动作。它是一种互联网应用程序的API设计理念：**URL定位资源，用HTTP描述操作** 

+ GET⽤来获取资源 
+ POST⽤来新建资源 
+ PUT⽤来更新资源 
+ DELETE⽤来删除资源。 

只要API程序遵循了REST⻛格，那就可以称其为RESTful API。⽬前在前后端分离的架构中，前后端基本都是通 过RESTful API来进⾏交互。



下面用Gin实现一下

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

func main(){
   r:=gin.Default()
   r.GET("/",func(c *gin.Context){
      c.JSON(http.StatusOK,gin.H{
         "message":"GET",
      })
   })

   r.POST("/post",func(c *gin.Context){
      c.JSON(http.StatusOK,gin.H{
         "message":"POST",
      })
   })

   r.PUT("/put",func(c *gin.Context){
      c.JSON(http.StatusOK,gin.H{
         "message":"PUT",
      })
   })

   r.DELETE("/delete",func(c *gin.Context){
      c.JSON(http.StatusOK,gin.H{
         "message":"DELETE",
      })
   })

   r.Run()
}
```

![image-20211111220016748](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211111220016748.png)



## 2. API参数问题

我们可以通过Context的Param方法来获取API参数

先来看一段demo代码

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
   "strings"
)

func main(){
   r:=gin.Default()
   r.GET("/user/:name/*action",func(c *gin.Context){
      name:=c.Param("name")
      action:=c.Param("action")
      // strings.Trim()返回去除所有包含cutset之后的结果
      action=strings.Trim(action,"/")
      c.String(http.StatusOK,"name:"+name +"\naction:"+ action)
   })
   r.Run()
}
```

![image-20211111223143022](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211111223143022.png)

但是不知道你们有没有疑惑，url里面设置API参数的地方使用:和\*是什么意思，当时我也有这种困惑，但是所有的文档都没有解释这个问题，所以我就自己看了源码来解释一下。

首先点击c.Param进去context.go，然后看到

![image-20211111225050391](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211111225050391.png)

可以看到这个其实就是根据键去找对应的值，继续点击ByName()方法，来到tree.go

![image-20211111231210385](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211111231210385.png)

`ps.Get()`用到的就是上面的Get()函数，遍历ps参数切片找对应的name，当name相同时返回对应的值和true，否则返回false；

![image-20211111231742864](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211111231742864.png)

再往上看，是Params参数切片的定义以及Param结构体的声明，很明显的键值对，根据上面的注释也能看懂Param是一个单一的URL参数，包含一个键和一个值。

到了最上面，看到了我们疑惑的地方，这里定义了:和\*，不过根据名字根本不知道这两个byte切片是干嘛的。**很不理解其他地方的注释都很详细，为什么这两个最令人困惑的地方定义居然一点注释都没有！！！**

既然这样我们继续深挖，tree.go定义了这两个切片，后面肯定有使用到的地方，从使用的地方说不定能看出它们的作用。果不其然，继续往后翻可以看到一个查找通配符的函数

![image-20211111232709082](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211111232709082.png)

上面的注释，**查找一个通配符字段并且检查无效字符名称，没有找到通配符就返回-1**

这其实还是没解决我们这个问题，在看函数内部的注释，通配符以:开始（参数）或者\*开始（获取全部）

到这里我们就把这个困惑解决了，**当以:开头的通配符只匹配一个参数，以\*开头的就把后面的所有内容全部匹配。**

知道了规则之后我们再来写点测试用例试试

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
   "strings"
)

func main(){
   r:=gin.Default()
   r.GET("/user/:name/:xxx/*action",func(c *gin.Context){
      name:=c.Param("name")
      xxx:=c.Param("xxx")
      action:=c.Param("action")
      // strings.Trim()返回去除所有包含cutset之后的结果
      action=strings.Trim(action,"/")
      c.String(http.StatusOK,"name:"+name+"\nxxx:"+xxx+"\naction:"+ action)
   })
   r.Run()
}
```

![image-20211111234053058](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211111234053058.png)

如果:出现在\*之后呢？

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
   "strings"
)

func main(){
   r:=gin.Default()
   r.GET("/user/:name/*action/:xxx",func(c *gin.Context){
      name:=c.Param("name")
      xxx:=c.Param("xxx")
      action:=c.Param("action")
      // strings.Trim()返回去除所有包含cutset之后的结果
      action=strings.Trim(action,"/")
      c.String(http.StatusOK,"name:"+name+"\nxxx:"+xxx+"\naction:"+ action)
   })
   r.Run()
}
```

![image-20211111234227230](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211111234227230.png)

编译器果然报错了，catch-all通配符只能在路径的最后，其实这个在tree.go中设计者就考虑到了这一点，所以如果不熟悉了解这两个通配符就可能写出的代码闹出笑话。

![image-20211111234514315](https://gitee.com/shixiaojiejiela_admin/pics/raw/master/upic/image-20211111234514315.png)

