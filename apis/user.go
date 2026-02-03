package apis

import (
	"gin-online-chat-backend/commons"
	"gin-online-chat-backend/commons/utils"
	"gin-online-chat-backend/services/user"

	"github.com/gin-gonic/gin"
)

func SaveUserGroup(c *gin.Context) {
	var req user.SaveUserGroupModel
	if err := utils.BindJsonToStruct(c, &req); err != nil {
		commons.Fail(c, commons.CodeParamError, "参数错误: "+err.Error())
		return
	}
	err := user.SaveUserGroup(c, &req)
	if err != nil {
		commons.Fail(c, commons.CodeServerError, err.Error())
		return
	}
	commons.Success(c, "新增分组成功", nil)
}
