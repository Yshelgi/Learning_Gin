package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//设置404 NOT FOUND

func main() {
	r := gin.New()
	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "shelgi")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	// 当访问到不知名路由
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 NOT FOUND")
	})
	r.Run()
}
