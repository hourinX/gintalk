package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "用户注册成功"})
		})
		userGroup.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
		})
	}
}
