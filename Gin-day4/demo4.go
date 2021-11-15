package main

import (
	"Learn_Gin/Gin-day4/routers"
	"github.com/gin-gonic/gin"
)

//多文件路由注册

func main() {
	r := gin.Default()
	routers.Test1(r)
	routers.Test2(r)
	r.Run()
}
