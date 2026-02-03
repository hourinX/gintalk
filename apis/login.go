package apis

import (
	"gin-online-chat-backend/commons"
	"gin-online-chat-backend/commons/utils"
	"gin-online-chat-backend/services/login"

	"github.com/gin-gonic/gin"
)

func CryptoPk(c *gin.Context) {
	var req login.PublicKeyModel
	if err := utils.BindQueryStrict(c, &req); err != nil {
		commons.Fail(c, commons.CodeParamError, "参数错误: "+err.Error())
		return
	}
	data, err := login.GetPublicKey(c, &req)
	if err != nil {
		commons.Fail(c, commons.CodeServerError, err.Error())
		return
	}
	commons.Success(c, "登录成功", data)
}

func Login(c *gin.Context) {
	var req login.UserLoginModel
	if err := utils.BindJsonToStruct(c, &req); err != nil {
		commons.Fail(c, commons.CodeParamError, "参数错误: "+err.Error())
		return
	}
	data, err := login.UserLogin(c, &req)
	if err != nil {
		commons.Fail(c, commons.CodeServerError, err.Error())
		return
	}
	commons.Success(c, "登录成功", data)
}

func Register(c *gin.Context) {
	var req login.UserRegisterModel
	if err := utils.BindJsonToStruct(c, &req); err != nil {
		commons.Fail(c, commons.CodeParamError, "参数错误: "+err.Error())
		return
	}
	code, err := login.UserRegister(&req)
	if err != nil {
		commons.Fail(c, commons.CodeServerError, err.Error())
		return
	}
	commons.Success(c, "注册成功", code)
}

func Captcha(c *gin.Context) {
	//data := user.UserCaptcha()

}
