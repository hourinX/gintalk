package apis

import (
	"gin-online-chat-backend/commons"
	"gin-online-chat-backend/commons/utils"
	"gin-online-chat-backend/services/system"

	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	var req system.SystemRefreshModel
	if err := utils.BindJsonToStruct(c, &req); err != nil {
		commons.Fail(c, commons.CodeParamError, "参数错误: "+err.Error())
		return
	}
	data, err := system.RefreshToken(&req)
	if err != nil {
		commons.Fail(c, commons.CodeServerError, err.Error())
		return
	}
	commons.Success(c, "刷新令牌成功", data)
}
