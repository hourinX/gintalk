package system

import (
	"errors"
	"gin-online-chat-backend/commons/utils"
	"gin-online-chat-backend/systems"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func RefreshToken(model *SystemRefreshModel) (*ReadSystemRefreshModel, error) {
	if model.RefreshToken == "" {
		return nil, errors.New("刷新令牌不能为空")
	}
	jwtSecret := []byte(systems.GetConfig().JWT.Secret)
	token, err := jwt.Parse(model.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("无效的刷新令牌")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("无法解析令牌")
	}

	if exp, ok := claims["exp"].(float64); ok && int64(exp) < time.Now().Unix() {
		return nil, errors.New("刷新令牌已过期")
	}

	userId, _ := claims["id"].(string)
	accessToken, expireTime, err := utils.GenerateToken(userId)
	if err != nil {
		return nil, err
	}

	data := &ReadSystemRefreshModel{
		AccessToken: accessToken,
		ExpireTime:  expireTime,
	}
	return data, nil
}
