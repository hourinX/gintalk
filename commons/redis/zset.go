package redis

import (
	"gin-online-chat-backend/systems"

	"github.com/redis/go-redis/v9"
)

func ZAdd(key, member string, score float64) error {
	return systems.RedisClient.ZAdd(systems.RedisCtx, key, redis.Z{Score: score, Member: member}).Err()
}

func ZTop(key string, k int64) ([]redis.Z, error) {
	return systems.RedisClient.ZRevRangeWithScores(systems.RedisCtx, key, 0, k-1).Result()
}

func ZIncr(key, member string, incr float64) (float64, error) {
	return systems.RedisClient.ZIncrBy(systems.RedisCtx, key, incr, member).Result()
}
