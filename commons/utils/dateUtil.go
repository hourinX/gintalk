package utils

import (
	"gin-online-chat-backend/commons"
	"time"
)

func ParseDate(t time.Time, f string) string {
	if f == "" {
		f = commons.YYYYMMDDHHMMSS
	}
	return t.Format(f)
}
