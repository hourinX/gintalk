package user

import (
	"errors"
	"gin-online-chat-backend/commons/utils"
	"gin-online-chat-backend/daos"
	"gin-online-chat-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(model *UserRegisterModel) error {
	if model.UserName == "" {
		return errors.New("用户名不能为空")
	}
	if model.Password == "" {
		return errors.New("密码不能为空")
	}

	if u, _ := daos.GetUserByCondition(&models.UserWhere{UserName: model.UserName}, "id"); u.Id != "" {
		return errors.New("用户名已存在")
	}
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &models.User{
		UserName: model.UserName,
		Password: string(hashedPwd),
		Avatar:   model.Avatar,
		Phone:    model.Phone,
		Email:    model.Email,
		Gender:   model.Gender,
	}
	err = daos.RegisterUser(user)
	if err != nil {
		return err
	}
	return nil
}

func UserLogin(model *UserLoginModel) (*ReadUserLoginModel, error) {
	u, err := daos.GetUserByCondition(&models.UserWhere{
		UserCode: model.UserCode,
	}, "id,user_code,user_name,password,phone,email,avatar,gender,is_frozen")
	if err != nil && (u).Id != "" {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte((u).Password), []byte(model.Password)); err != nil {
		return nil, err
	}
	token, err := utils.GenerateToken(u.Id)
	if err != nil {
		return nil, err
	}
	data := &ReadUserLoginModel{
		Id:       u.Id,
		UserName: u.UserName,
		UserCode: u.UserCode,
		Avatar:   u.Avatar,
		Phone:    u.Phone,
		Email:    u.Email,
		Gender:   u.Gender,
		IsFrozen: u.IsFrozen,
		Token:    token,
	}
	return data, nil
}

func UserCaptcha() (string, error) {
	return "", nil
}
