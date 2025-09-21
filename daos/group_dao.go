package daos

import (
	"gin-online-chat-backend/commons"
	"gin-online-chat-backend/commons/utils"
	"gin-online-chat-backend/models"
	"gin-online-chat-backend/systems"
	"time"

	"gorm.io/gorm"
)

func InsertUserContact(tx *gorm.DB, user *models.UserContact) error {
	if tx == nil {
		tx = systems.DB
	}
	id, err := utils.GenerateNumericID(commons.MaxIDLength)
	if err != nil {
		return err
	}
	user.Id = id
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	user.Status = commons.StatusUp
	return tx.Table(commons.TableImUserContacts).Create(user).Error
}

func InsertUserGroup(tx *gorm.DB, group *models.UserGroup) error {
	if tx == nil {
		tx = systems.DB
	}
	id, err := utils.GenerateNumericID(commons.MaxIDLength)
	if err != nil {
		return err
	}
	curTime := time.Now()
	group.Id = id
	group.CreateTime = curTime
	group.UpdateTime = curTime
	return tx.Table(commons.TableImUserGroups).Create(group).Error
}
