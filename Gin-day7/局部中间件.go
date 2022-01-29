package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func MiddleWare1() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println(c.Request.URL.Path + "局部中间件开始")
		c.Set("request", "中间件")
		c.Next()
		status := c.Writer.Status()
		fmt.Println("局部中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("耗时:", t2)
	}
}

func main() {
	r := gin.Default()

	r.GET("test", MiddleWare1(), func(c *gin.Context) {
		req, _ := c.Get("request")
		fmt.Println("request:", req)
		c.JSON(200, gin.H{"request": req})
	})

	r.Run()
}
