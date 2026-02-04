package router

import (
	"gin-online-chat-backend/apis"
	"gin-online-chat-backend/commons/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeLWsRouter(r *gin.Engine) {
	wsGroup := r.Group("")
	wsGroup.Use(middleware.JWTWsMiddleware())
	{
		wsGroup.GET("/ws", apis.WsHandler)
	}
}
