package test

import (
	"gin-online-chat-backend/commons/redis"
	"testing"
)

func RedisTestFunc(t *testing.T) {
	redis.Set("test_key", "test_value", 0)

	var value string
	err := redis.Get("test_key", &value)
	if err != nil {
		t.Errorf("Failed to get value from Redis: %v", err)
	} else if value != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", value)
	}
}
