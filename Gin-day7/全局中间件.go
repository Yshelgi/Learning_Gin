package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("全局中间件开始")
		c.Set("request", "中间件")
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("耗时:", t2)
	}
}

func main() {
	r := gin.Default()
	r.Use(MiddleWare())

	r.GET("test", func(c *gin.Context) {
		req, _ := c.Get("request")
		fmt.Println("request:", req)
		c.JSON(200, gin.H{"request": req})
	})

	r.Run()
}
