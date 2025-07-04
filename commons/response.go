package commons

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  msg,
		"data": data,
	})
}

func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
