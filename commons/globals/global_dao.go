package globals

import (
	"gin-online-chat-backend/commons"
	"gin-online-chat-backend/systems"
)

func IsCodeExistsInUser(code string) (bool, error) {
	var count int64
	systems.DB.Table(commons.TableImUser).Select("user_code").Where("user_code = ?", code).Count(&count)
	if count > 0 {
		return true, nil
	}
	return false, nil
}
