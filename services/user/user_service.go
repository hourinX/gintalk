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
	if model.Phone == "" {
		return errors.New("手机号不能为空")
	}
	if u, _ := daos.GetUserByCondition(&models.UserWhere{Phone: model.Phone}, "id"); u.Id != "" {
		return errors.New("当前手机号已被注册")
	}
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &models.User{
		UserName: model.UserName,
		Password: string(hashedPwd),
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
	if model.UserCode == "" {
		return nil, errors.New("用户账号不能为空")
	}
	if model.Password == "" {
		return nil, errors.New("密码不能为空")
	}
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
