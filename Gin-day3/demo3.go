package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//多文件上传

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s"), err.Error())
		}
		files := form.File["files"]
		for _, file := range files {
			if err := c.SaveUploadedFile(file, "./res/"+file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("upload %d files", len(files)))
	})

	r.Run()
}
