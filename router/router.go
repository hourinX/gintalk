package router

import (
	"gin-online-chat-backend/commons"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// r := gin.New()
	// r.Use(gin.Recovery())
	// r.Use(logger.RequestLogger())
	r.Use(commons.Cors())

	InitializeSystemRouter(r)
	InitializeLoginRouter(r)
	InitializeUserRouter(r)
	InitializeLogRouter(r)
	InitializeLWsRouter(r)
	// InitializeGroupRouter(r)

	return r
}
