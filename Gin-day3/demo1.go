package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//表单参数

func main() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		// 键名和页面属性名对应
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,types:%s", username, password, types))
	})
	r.Run()
}
