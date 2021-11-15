package demo

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine) {
	e.GET("/id", idHandler)
	e.GET("/comment", commentHandler)
}
