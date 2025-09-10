package router

import (
	"gin-online-chat-backend/apis"
	"gin-online-chat-backend/commons/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeLogRouter(r *gin.Engine) {
	logGroup := r.Group("/log")
	logGroup.Use(middleware.JWTAuthMiddleware())
	{
		logGroup.GET("/getLoginLogList", apis.GetLoginLogList)
	}
}
