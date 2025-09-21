package apis

import (
	"gin-online-chat-backend/commons"
	"gin-online-chat-backend/commons/utils"
	"gin-online-chat-backend/services/group"

	"github.com/gin-gonic/gin"
)

func GetUserGroups(c *gin.Context) {
	var req group.UserGroupModal
	if err := utils.BindQueryStrict(c, &req); err != nil {
		commons.Fail(c, commons.CodeParamError, "参数错误: "+err.Error())
		return
	}
	data, err := group.GetUserGroups(&req)
	if err != nil {
		commons.Fail(c, commons.CodeServerError, err.Error())
		return
	}
	commons.Success(c, "查询成功", data)
}
