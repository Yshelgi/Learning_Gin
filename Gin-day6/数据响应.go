package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {
	r := gin.Default()

	// JSON格式
	r.GET("/JSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
	})

	// 结构体格式
	r.GET("/Struct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}

		msg.Name = "test"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200, msg)
	})

	// XML
	r.GET("/XML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "abc"})

	})

	// YAML
	r.GET("/YAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "test"})
	})

	// protobuf
	r.GET("/Protobuf", func(c *gin.Context) {
		resp := []int64{int64(10), int64(2)}
		label := "label"
		data := &protoexample.Test{
			Label: &label,
			Reps:  resp,
		}
		c.ProtoBuf(200, data)
	})

	r.Run(":8888")
}
