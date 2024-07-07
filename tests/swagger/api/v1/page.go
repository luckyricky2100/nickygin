package v1

import "github.com/gin-gonic/gin"

//
// @Summary GetClass
// @Description GetPage
// @Accept  json
// @Produce  json
// @Param  some_msg    path    string     true        "Some Msg"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /api/v1/{some_msg} [get]
func GetPage(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"first": "page:",
	})
}
