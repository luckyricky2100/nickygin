package api

import (
	"github.com/gin-gonic/gin"
	"nickygin.com/global"
	"nickygin.com/pkg/app"
	"nickygin.com/pkg/errcode"
	"nickygin.com/recitewords/internal/service"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	auth, err := svc.GetAuth(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	claims := app.Claims{Uid: param.AppKey, NickName: auth.Nickname}
	token, err := app.GenerateToken(claims)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}

func Register(c *gin.Context) {
	param := service.RegisterRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.Register(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.Rigister err: %v", err)
		response.ToErrorResponse(errcode.RigisterFail.WithDetails(err.Error()))
		return
	} else {
		response.ToResponse(gin.H{
			"message": "注册成功",
		})
	}
}
