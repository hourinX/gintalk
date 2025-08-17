package router

import (
	"gin-online-chat-backend/apis"

	"github.com/gin-gonic/gin"
)

func InitializeSystemRouter(r *gin.Engine) {
	systemGroup := r.Group("/system")
	{
		systemGroup.POST("/refresh", apis.RefreshToken)
	}
}
