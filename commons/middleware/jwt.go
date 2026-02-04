package middleware

import (
	"gin-online-chat-backend/systems"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		jwtScrect := []byte(systems.GetConfig().JWT.Secret)
		issuer := systems.GetConfig().JWT.Issuer

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtScrect, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "cannot resolved claims"})
			c.Abort()
			return
		}

		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "access_token has been expired"})
				c.Abort()
				return
			}
		}

		if iss, ok := claims["iss"].(string); !ok || iss != issuer {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token issuer"})
			c.Abort()
			return
		}

		if uid, ok := claims["id"].(string); ok {
			c.Set("userId", uid)
			return
		}

		c.Next()
	}
}

func JWTWsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			tokenStr = c.Query("token")
			if tokenStr != "" {
				tokenStr = "Bearer " + tokenStr
			}
		}

		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		jwtScrect := []byte(systems.GetConfig().JWT.Secret)
		issuer := systems.GetConfig().JWT.Issuer

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtScrect, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "cannot resolved claims"})
			c.Abort()
			return
		}

		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "access_token has been expired"})
				c.Abort()
				return
			}
		}

		if iss, ok := claims["iss"].(string); !ok || iss != issuer {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token issuer"})
			c.Abort()
			return
		}

		if uid, ok := claims["id"].(string); ok {
			c.Set("userId", uid)
			return
		}

		c.Next()
	}
}
