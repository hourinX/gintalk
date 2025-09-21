package router

import (
	"gin-online-chat-backend/apis"
	"gin-online-chat-backend/commons/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeGroupRouter(r *gin.Engine) {
	groups := r.Group("/group")
	groups.Use(middleware.JWTAuthMiddleware())
	{
		groups.POST("/getUserGroups", apis.GetUserGroups)
	}
}
