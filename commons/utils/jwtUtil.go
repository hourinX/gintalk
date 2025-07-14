package utils

import (
	"fmt"
	"gin-online-chat-backend/systems"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(userId string) (string, error) {
	fmt.Println(systems.GetConfig().JWT.Secret)
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Hour * time.Duration(systems.GetConfig().JWT.ExpireHours)).Unix(),
		"iss": systems.GetConfig().JWT.Issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte(systems.GetConfig().JWT.Secret)
	return token.SignedString(jwtSecret)
}
