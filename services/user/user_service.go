package user

import (
	"errors"
	"fmt"
	"gin-online-chat-backend/commons"
	"gin-online-chat-backend/commons/utils"
	"gin-online-chat-backend/daos"
	"gin-online-chat-backend/models"
	"gin-online-chat-backend/systems"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(model *UserRegisterModel) (string, error) {
	if model.UserName == "" {
		return "", errors.New("用户名不能为空")
	}
	if model.Password == "" {
		return "", errors.New("密码不能为空")
	}
	if model.Phone == "" {
		return "", errors.New("手机号不能为空")
	}
	if u, _ := daos.GetUserByCondition(&models.UserWhere{Phone: model.Phone}, "id"); u.Id != "" {
		return "", errors.New("当前手机号已被注册")
	}
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &models.User{
		UserName: model.UserName,
		Password: string(hashedPwd),
		Phone:    model.Phone,
		Email:    model.Email,
		Gender:   model.Gender,
	}
	code, err := daos.RegisterUser(user)
	if err != nil {
		return "", err
	}
	return code, nil
}

func UserLogin(c *gin.Context, model *UserLoginModel) (*ReadUserLoginModel, error) {
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
		return nil, fmt.Errorf("用户名或密码错误,请重试")
	}
	accessToken, expireTime, err := utils.GenerateToken(u.Id)
	if err != nil {
		return nil, err
	}
	refreshToken, _ := utils.GenerateRefreshToken(u.Id)

	ip := c.ClientIP()
	userAgent := c.Request.UserAgent()
	ua := user_agent.New(userAgent)
	browserName, browserVersion := ua.Browser()
	os := ua.OS()
	deviceInfo := fmt.Sprintf("%s %s / %s", browserName, browserVersion, os)

	err = daos.InsertLoginLog(nil, &models.LoginLogs{
		UserId:      u.UserCode,
		UserName:    u.UserName,
		LoginTime:   time.Now(),
		IpAddress:   ip,
		DeviceInfo:  deviceInfo,
		LoginStatus: "success",
		LoginMethod: "password",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to insert login log: %w", err)
	}

	data := &ReadUserLoginModel{
		Id:           u.Id,
		UserName:     u.UserName,
		UserCode:     u.UserCode,
		Avatar:       u.Avatar,
		Phone:        u.Phone,
		Email:        u.Email,
		Gender:       u.Gender,
		IsFrozen:     u.IsFrozen,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpireTime:   expireTime,
	}

	err = systems.Set(fmt.Sprintf("user:%s", u.UserCode), data, 24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to cache user data: %w", err)
	}
	return data, nil
}

func UserCaptcha() (string, error) {
	return "", nil
}

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
