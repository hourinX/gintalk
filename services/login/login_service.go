package login

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"gin-online-chat-backend/commons/redis"
	"gin-online-chat-backend/commons/utils"
	"gin-online-chat-backend/daos"
	"gin-online-chat-backend/models"
	"gin-online-chat-backend/systems"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(c *gin.Context, model *UserLoginModel) (*ReadUserLoginModel, error) {
	if model.UserCode == "" {
		return nil, errors.New("用户账号不能为空")
	}
	if model.Password == "" {
		return nil, errors.New("密码不能为空")
	}

	rsaCipherAESKey, err := base64.StdEncoding.DecodeString(model.EncryptedAESKey)
	if err != nil {
		return nil, errors.New("AES Key base64 解码失败")
	}
	aesKey, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, systems.PrivateKey, rsaCipherAESKey, nil)
	if err != nil {
		return nil, errors.New("AES Key 解密失败")
	}
	if len(aesKey) != 16 && len(aesKey) != 24 && len(aesKey) != 32 {
		return nil, fmt.Errorf("AES key 长度不正确: %d", len(aesKey))
	}
	plainPasswordBytes, err := utils.AESDecrypt(model.Password, aesKey)
	if err != nil {
		return nil, errors.New("密码 AES 解密失败" + err.Error())
	}
	plainPassword := string(plainPasswordBytes)

	u, err := daos.GetUserByCondition(&models.UserWhere{
		UserCode: model.UserCode,
	}, "id,user_code,user_name,password,phone,email,avatar,gender,is_frozen")
	if err != nil && (u).Id != "" {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte((u).Password), []byte(string(plainPassword))); err != nil {
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
		LoginMethod: "账密登录",
	})
	if err != nil {
		return nil, fmt.Errorf("新增登录日志失败: %w", err)
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

	err = redis.Set(fmt.Sprintf("user:%s", u.UserCode), data, 24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("缓存用户信息失败: %w", err)
	}
	return data, nil
}

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

func UserCaptcha() (string, error) {
	return "", nil
}

func GetPublicKey(c *gin.Context, model *PublicKeyModel) (*ReadPublicRsaKeyModel, error) {
	aesKey := make([]byte, 32)
	if _, err := rand.Read(aesKey); err != nil {
		return nil, err
	}
	encAesKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, systems.PublicKey, []byte(aesKey), nil)
	if err != nil {
		return nil, err
	}

	result := &ReadPublicRsaKeyModel{
		AesKey:      base64.StdEncoding.EncodeToString(encAesKey),
		PlainAesKey: base64.StdEncoding.EncodeToString(aesKey),
	}
	return result, nil
}
