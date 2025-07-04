package daos

import (
	"gin-online-chat-backend/commons/utils"
	"gin-online-chat-backend/models"
	"gin-online-chat-backend/systems"
)

func RegisterUser(user *models.User) error {
	id, err := utils.GenerateNumericID()
	if err != nil {
		return err
	}
	user.Id = id
	return systems.DB.Create(user).Error
}

func GetUserByUserName(username string) (*models.User, error) {
	var user models.User
	err := systems.DB.Where("user_name = ?", username).First(&user).Error
	return &user, err
}
