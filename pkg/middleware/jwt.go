package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"nickygin.com/global"
	"nickygin.com/pkg/app"
	"nickygin.com/pkg/errcode"
	"nickygin.com/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)

		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			ecode = errcode.InvalidParams.WithDetails("please set token first.")
		} else {
			claims, err := app.ParseToken(token)
			if err != nil {
				switch err {
				case jwt.ErrTokenExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			} else {
				uid, err := util.AESDecrypt(claims.Uid, global.JWTSetting.Secret2)
				if err != nil {
					response := app.NewResponse(c)
					response.ToErrorResponse(errcode.UnauthorizedTokenError)
					return
				}
				c.Set("nickname", claims.NickName)
				c.Set("uid", uid)
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return

		}

		c.Next()
	}
}
