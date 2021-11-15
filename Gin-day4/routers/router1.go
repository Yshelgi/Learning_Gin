package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func testhandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"test": "这是路由注册的测试",
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/test", testhandler)
	return r
}
