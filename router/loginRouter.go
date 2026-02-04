package router

import (
	"gin-online-chat-backend/apis"

	"github.com/gin-gonic/gin"
)

func InitializeLoginRouter(r *gin.Engine) {
	loginGroup := r.Group("/auth")
	{
		loginGroup.GET("/cryptoPk", apis.CryptoPk)
		loginGroup.POST("/login", apis.Login)
		loginGroup.POST("/register", apis.Register)
		loginGroup.GET("/captcha", apis.Captcha)
	}
}
