package v1

import (
	"github.com/gin-gonic/gin"
	"nickygin.com/global"
	"nickygin.com/pkg/app"
	"nickygin.com/pkg/errcode"
	"nickygin.com/recitewords/internal/service"
)

type UserClass struct {
}

func NewUserClass() UserClass {
	return UserClass{}
}

func (*UserClass) List(c *gin.Context) {
	param := service.UserKnowledgeClassRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	param.UserId = c.GetString("uid")
	classes, err := svc.GetClasses(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserClassListFail)
		return
	}
	response.ToResponseList(classes, len(classes))
}

// @Summary GetClass
// @Description GetAllClasses
// @Accept  mpfd
// @Produce  json
// @Param  title    formData    string     false        "string"
// @Param  parent_id    formData    int     false        "int"
// @Param  enable_to_memory    formData    bool     false        "bool"
// @Param  token    header    string     false        "string"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} errcode.Error "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /api/v1/classes [post]
func (*UserClass) Create(c *gin.Context) {
	param := service.CreateUserClass{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	param.UserId = c.GetString("uid")
	svc := service.New(c.Request.Context())
	class, err := svc.CreateClass(&param)
	if err != nil && (class == nil || class.ID == 0) {
		global.Logger.Errorf(c, "svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateUserClassFail)
		return
	}

	response.ToResponse(gin.H{"message": "创建成功"})
}

// // @Summary 更新标签
// // @Produce  json
// // @Param id path int true "标签ID"
// // @Param name body string false "标签名称" minlength(3) maxlength(100)
// // @Param state body int false "状态" Enums(0, 1) default(1)
// // @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// // @Success 200 {array} model.Tag "成功"
// // @Failure 400 {object} errcode.Error "请求错误"
// // @Failure 500 {object} errcode.Error "内部错误"
// // @Router /api/v1/tags/{id} [put]
// func (t Tag) Update(c *gin.Context) {
// 	param := service.UpdateTagRequest{
// 		ID: convert.StrTo(c.Param("id")).MustUInt32(),
// 	}
// 	response := app.NewResponse(c)
// 	valid, errs := app.BindAndValid(c, &param)
// 	if !valid {
// 		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
// 		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
// 		return
// 	}

// 	svc := service.New(c.Request.Context())
// 	err := svc.UpdateTag(&param)
// 	if err != nil {
// 		global.Logger.Errorf(c, "svc.UpdateTag err: %v", err)
// 		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
// 		return
// 	}

// 	response.ToResponse(gin.H{})
// 	return
// }

// // @Summary 删除标签
// // @Produce  json
// // @Param id path int true "标签ID"
// // @Success 200 {string} string "成功"
// // @Failure 400 {object} errcode.Error "请求错误"
// // @Failure 500 {object} errcode.Error "内部错误"
// // @Router /api/v1/tags/{id} [delete]
// func (t Tag) Delete(c *gin.Context) {
// 	param := service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
// 	response := app.NewResponse(c)
// 	valid, errs := app.BindAndValid(c, &param)
// 	if !valid {
// 		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
// 		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
// 		return
// 	}

// 	svc := service.New(c.Request.Context())
// 	err := svc.DeleteTag(&param)
// 	if err != nil {
// 		global.Logger.Errorf(c, "svc.DeleteTag err: %v", err)
// 		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
// 		return
// 	}

// 	response.ToResponse(gin.H{})
// 	return
// }
