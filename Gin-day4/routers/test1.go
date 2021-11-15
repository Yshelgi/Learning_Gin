package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func test1handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"test": "这是路由1注册的测试",
	})
}

func Test1(e *gin.Engine) {
	e.GET("/test1", test1handler)
}
