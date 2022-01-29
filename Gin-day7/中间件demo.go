package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func myMiddle(c *gin.Context) {
	s := time.Now()
	c.Next()
	e := time.Since(s)
	fmt.Println("耗时:", e)
}

func handler1(c *gin.Context) {
	time.Sleep(1 * time.Second)
	c.JSON(200, gin.H{"handler1": "完成"})
}

func handler2(c *gin.Context) {
	time.Sleep(2 * time.Second)
	c.JSON(200, gin.H{"handler2": "完成"})
}

func main() {
	f, _ := os.Create("demo.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r.Use(myMiddle)
	Group := r.Group("/test")
	{
		Group.GET("/test1", handler1)
		Group.GET("/test2", handler2)
	}
	r.Run()
}
