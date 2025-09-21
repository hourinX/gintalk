package daos

import (
	"gin-online-chat-backend/commons"
	"gin-online-chat-backend/commons/utils"
	"gin-online-chat-backend/models"
	"gin-online-chat-backend/systems"
	"time"
)

func RegisterUser(user *models.User) (string, error) {
	id, err := utils.GenerateNumericID(commons.MaxIDLength)
	if err != nil {
		return "", err
	}
	user.Id = id
	user.IsFrozen = 1
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	if user.UserCode == "" {
		code, err := utils.GenerateUniqueUserCodeWithRetry()
		if err != nil {
			return "", err
		}
		user.UserCode = code
	}
	return user.UserCode, systems.DB.Table(commons.TableImUser).Create(user).Error
}

func GetUserByCondition(w *models.UserWhere, field string) (*models.User, error) {
	var user models.User
	where, args := parseUserWhere(w)
	err := systems.DB.Table(commons.TableImUser).Select(field).Where(where, args...).First(&user).Error
	return &user, err
}
