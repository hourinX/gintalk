package systems

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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

func Set(key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return RedisClient.Set(ctx, key, data, expiration).Err()
}

func Get(key string, result interface{}) error {
	data, err := RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

func Delete(key string) error {
	return RedisClient.Del(ctx, key).Err()
}
