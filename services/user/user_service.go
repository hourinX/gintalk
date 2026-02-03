package user

import (
	"errors"
	"gin-online-chat-backend/commons"
	"gin-online-chat-backend/daos"
	"gin-online-chat-backend/models"

	"github.com/gin-gonic/gin"
)

func SaveUserGroup(c *gin.Context, model *SaveUserGroupModel) error {
	if model.UserId == "" {
		return errors.New("用户ID不能为空")
	}
	if model.GroupName == "" {
		return errors.New("分组名称不能为空")
	}

	data := &models.UserGroup{
		UserId:    model.UserId,
		GroupName: model.GroupName,
		Type:      commons.GroupTypeCustom,
		Editable:  commons.EditablePermission,
	}

	err := daos.InsertUserGroup(nil, data)
	if err != nil {
		return err
	}

	return nil
}
