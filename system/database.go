package system

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

var DB *gorm.DB

func InitDatabase() error {
	//构建数据库链接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		GetConfig().MySQL.Username,
		GetConfig().MySQL.Password,
		GetConfig().MySQL.Host,
		strconv.Itoa(GetConfig().MySQL.Port),
		GetConfig().MySQL.Database,
		GetConfig().MySQL.Charset,
		GetConfig().MySQL.ParseTime,
		GetConfig().MySQL.Loc,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	fmt.Println("✅ 数据库链接成功")
	return nil
}

func CloseDatabase() {
	sqlDB, err := DB.DB()
	if err != nil {
		_ = sqlDB.Close()
	}
}
