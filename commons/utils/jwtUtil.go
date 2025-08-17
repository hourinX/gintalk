package utils

import (
	"errors"
	"gin-online-chat-backend/systems"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId string) (string, int64, error) {
	jwtSecret := []byte(systems.GetConfig().JWT.Secret)

	expireTime := time.Now().Add(time.Hour * time.Duration(systems.GetConfig().JWT.ExpireHours)).Unix()
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": expireTime,
		"iss": systems.GetConfig().JWT.Issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", 0, err
	}

	return signedToken, expireTime, nil
}

func GenerateRefreshToken(userId string) (string, int64) {
	jwtScrect := []byte(systems.GetConfig().JWT.Secret)

	expireTime := time.Now().Add(time.Hour * 24 * time.Duration(systems.GetConfig().JWT.RefreshDays)).Unix()
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": expireTime,
		"iss": systems.GetConfig().JWT.Issuer,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := refreshToken.SignedString(jwtScrect)
	if err != nil {
		return "", 0
	}

	return signedToken, expireTime
}

func RedreshAccessToken(refreshToken string) (string, int64, error) {
	jwtSecret := []byte(systems.GetConfig().JWT.Secret)

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return "", 0, errors.New("refresh_token are expired or invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", 0, errors.New("refresh_token claims are cannot be parsed")
	}

	expUnix, ok := claims["exp"].(float64)
	if !ok {
		return "", 0, errors.New("refresh_token exp is not a valid number")
	}
	if int64(expUnix) < time.Now().Unix() {
		return "", 0, errors.New("refresh_token are expired, please try to login again")
	}

	userId := claims["id"].(string)
	newAccessToken, newExpireTime, err := GenerateToken(userId)
	if err != nil {
		return "", 0, err
	}
	return newAccessToken, newExpireTime, nil
}
