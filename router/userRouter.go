package router

import (
	"gin-online-chat-backend/commons/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeUserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.Use(middleware.JWTAuthMiddleware())
	{
		// userGroup.GET("/profile", apis.GetUserProfile)
	}
}
