package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func idHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id_test": "这是id的测试",
	})
}

func commentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"comment_test": "这是comment的测试",
	})
}
