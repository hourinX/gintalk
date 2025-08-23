package apis

import (
	"gin-online-chat-backend/commons"
	"gin-online-chat-backend/commons/utils"
	"gin-online-chat-backend/services/user"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req user.UserRegisterModel
	if err := utils.BindJsonToStruct(c, &req); err != nil {
		commons.Fail(c, commons.CodeParamError, "参数错误: "+err.Error())
		return
	}
	code, err := user.UserRegister(&req)
	if err != nil {
		commons.Fail(c, commons.CodeServerError, err.Error())
		return
	}
	commons.Success(c, "注册成功", code)
}

func Login(c *gin.Context) {
	var req user.UserLoginModel
	if err := utils.BindJsonToStruct(c, &req); err != nil {
		commons.Fail(c, commons.CodeParamError, "参数错误: "+err.Error())
		return
	}
	data, err := user.UserLogin(c, &req)
	if err != nil {
		commons.Fail(c, commons.CodeServerError, err.Error())
		return
	}
	commons.Success(c, "登录成功", data)
}

func Captcha(c *gin.Context) {
	//data := user.UserCaptcha()

}
