package user

import (
	"errors"
	"gin-online-chat-backend/daos"
	"gin-online-chat-backend/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func UserRegister(model *UserRegisterModel) error {
	if u, _ := daos.GetUserByUserName(model.UserName); u.Id != "" {
		return errors.New("用户名已存在")
	}
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &models.User{
		UserName:  model.UserName,
		Password:  string(hashedPwd),
		Avatar:    model.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return daos.RegisterUser(user)
}

func UserLogin(model *UserLoginModel) (*models.User, error) {
	u, err := daos.GetUserByUserName(model.UserName)
	if err != nil {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(model.Password)); err != nil {
		return nil, err
	}
	return u, nil
}
