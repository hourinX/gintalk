package apis

import (
	"gin-online-chat-backend/commons"
	"gin-online-chat-backend/services/user"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req user.UserRegisterModel
	if err := c.ShouldBindJSON(&req); err != nil {
		commons.Fail(c, commons.CodeParamError, "参数错误: "+err.Error())
		return
	}
	err := user.UserRegister(&req)
	if err != nil {
		commons.Fail(c, commons.CodeServerError, err.Error())
		return
	}
	commons.Success(c, "注册成功", nil)
}

func Login(c *gin.Context) {
	var req user.UserLoginModel
	if err := c.ShouldBindJSON(&req); err != nil {
		commons.Fail(c, commons.CodeParamError, "参数错误: "+err.Error())
		return
	}
	token, err := user.UserLogin(&req)
	if err != nil {
		commons.Fail(c, commons.CodeServerError, err.Error())
		return
	}
	commons.Success(c, "登录成功", token)
}
