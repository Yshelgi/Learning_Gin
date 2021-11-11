package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/:xxx/*action", func(c *gin.Context) {
		name := c.Param("name")
		xxx := c.Param("xxx")
		action := c.Param("action")
		// strings.Trim()返回去除所有包含cutset之后的结果
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, "name:"+name+"\nxxx:"+xxx+"\naction:"+action)
	})
	r.Run()
}
