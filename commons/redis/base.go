package redis

import (
	"encoding/json"
	"gin-online-chat-backend/systems"
	"time"
)

func Set(key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return systems.RedisClient.Set(systems.RedisCtx, key, data, expiration).Err()
}

func Get(key string, result interface{}) error {
	data, err := systems.RedisClient.Get(systems.RedisCtx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

func Delete(key string) error {
	return systems.RedisClient.Del(systems.RedisCtx, key).Err()
}
