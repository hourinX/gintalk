package redis

import (
	"gin-online-chat-backend/systems"
	"time"
)

func TryLock(lockKey string, ttl time.Duration) (bool, error) {
	return systems.RedisClient.SetNX(systems.RedisCtx, lockKey, 1, ttl).Result()
}

func Unlock(lockKey string) error {
	return systems.RedisClient.Del(systems.RedisCtx, lockKey).Err()
}
