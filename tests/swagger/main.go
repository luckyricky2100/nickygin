package main

import (
	v1 "nickygin.com/tests/swagger/api/v1"
	_ "nickygin.com/tests/swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()

	// ... Define your routes here ...

	// Setup Swagger UI

	r.GET("/index", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"first": "page",
		})
	})
	r.GET("/getpage/:some_msg", v1.GetPage)
	//docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080
}
