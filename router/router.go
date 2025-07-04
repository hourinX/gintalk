package router

import (
	"gin-online-chat-backend/apis"
	"gin-online-chat-backend/commons"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(commons.Cors())
	apiUserGroup := r.Group("user")
	{
		apiUserGroup.POST("register", apis.Register)
		apiUserGroup.POST("login", apis.Login)
	}
	//apis.Register(r)
	return r
}
