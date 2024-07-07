package main

import (
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required,min=6"`
	Password string `form:"password" binding:"required"`
}

func TestValidation() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		// 你可以使用显式绑定声明绑定 multipart form：
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定：
		form := LoginForm{}
		err := c.ShouldBind(&form)
		// 在这种情况下，将自动选择合适的绑定
		if err == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(403, gin.H{"status": err.Error()})
		}
	})
	router.Run(":8080")
}
