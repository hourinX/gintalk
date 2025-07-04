package system

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis() error {
	conf := GetConfig().Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.DB,
	})
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}
	fmt.Println("✅ Redis链接成功")
	return nil
}

func CloseRedis() {
	if RedisClient != nil {
		_ = RedisClient.Close()
	}
}
