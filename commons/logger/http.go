package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 先执行后续逻辑
		c.Next()

		cost := time.Since(start)

		status := c.Writer.Status()
		success := status < 400

		zap.L().Info("http request",
			zap.String("type", "http"),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", status),
			zap.Bool("success", success),
			zap.Int64("cost_us", cost.Microseconds()),
			zap.String("ip", c.ClientIP()),
			zap.String("user_id", c.GetString("userId")),
		)
	}
}
