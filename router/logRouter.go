package router

import (
	"gin-online-chat-backend/apis"

	"github.com/gin-gonic/gin"
)

func InitializeLogRouter(r *gin.Engine) {
	logGroup := r.Group("/log")
	{
		logGroup.GET("/getLoginLogList", apis.GetLoginLogList)
	}
}
