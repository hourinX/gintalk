package test

import (
	"fmt"
	"gin-online-chat-backend/commons"
	"gin-online-chat-backend/commons/utils"
	"gin-online-chat-backend/models"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestUserInsert(t *testing.T) {
	var DB *gorm.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		"root",
		"Szjm270017!",
		"127.0.0.1",
		3306,
		"onlinechat",
		"utf8mb4",
		true,
		"Local",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	DB = db

	for i := 0; i < 1000; i++ {
		id, err := utils.GenerateNumericID(commons.MaxIDLength)
		if err != nil {
			return
		}
		user := models.User{
			Id:             id,
			IsFrozen:       1,
			UserCode:       fmt.Sprint("100000", i+1),
			UserName:       fmt.Sprintf("测试用户%04d", i+1),
			Password:       "123456",
			Phone:          fmt.Sprintf("1380000%04d", i+1),
			Avatar:         "",
			Gender:         1,
			Email:          fmt.Sprintf("test%04d@example.com", i+1),
			ImOnlineStatus: 1,
			CreateTime:     time.Now(),
			UpdateTime:     time.Now(),
		}

		DB.Table(commons.TableImUser).Create(&user)

		contactId, err := utils.GenerateNumericID(commons.MaxIDLength)
		if err != nil {
			return
		}
		users := &models.UserContact{
			Id:            contactId,
			UserId:        "17561290594990000005",
			ContactUserId: user.Id,
			Alias:         fmt.Sprintf("测试用户%04d", i+1),
			Status:        1,
			CreateTime:    time.Now(),
			UpdateTime:    time.Now(),
		}

		DB.Table(commons.TableImUserContacts).Create(users)
		if err != nil {
			t.Errorf("注册用户失败: %v", err)
			continue
		}

		time.Sleep(5 * time.Millisecond)
	}

}
