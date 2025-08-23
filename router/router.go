package router

import (
	"gin-online-chat-backend/commons"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(commons.Cors())

	InitializeSystemRouter(r)
	InitializeLoginRouter(r)
	InitializeUserRouter(r)
	InitializeLogRouter(r)

	return r
}
