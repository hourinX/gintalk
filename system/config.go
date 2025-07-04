package system

import (
	"fmt"
	"gin-online-chat-backend/models"
	"github.com/spf13/viper"
)

var config *models.Config

func LoadConfig(configPath string) error {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	var c models.Config
	if err := viper.Unmarshal(&c); err != nil {
		return err
	}
	config = &c
	fmt.Println("✅ 配置文件加载成功")
	return nil
}

func GetConfig() *models.Config {
	return config
}
