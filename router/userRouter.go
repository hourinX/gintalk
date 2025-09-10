package router

import (
	"gin-online-chat-backend/commons/middleware"
	"gin-online-chat-backend/apis"
	"github.com/gin-gonic/gin"
)

func InitializeUserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.Use(middleware.JWTAuthMiddleware())
	{
		userGroup.POST("/saveUserGroup", apis.SaveUserGroup)
	}
}
