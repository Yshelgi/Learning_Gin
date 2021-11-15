package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func test2handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"test": "这是路由2注册的测试",
	})
}

func Test2(e *gin.Engine) {
	e.GET("/test2", test2handler)
}
