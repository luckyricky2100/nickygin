package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestWeb() {
	r := gin.Default()

	// 提供 unicode 实体
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 提供字面字符
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK,
			gin.H{
				"message": "hey", "status": http.StatusOK})
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
