package dao

import (
	"fmt"
	"time"
)

const (
	_PREFIX = "QNDLD"
)

func Incr(ip string, ttl time.Duration) (int, error) {
	key := fmt.Sprintf("%v_%v", _PREFIX, ip)

	result, err := redisClient.Incr(key).Result()
	if err != nil {
		return 0, err
	}
	t, err := redisClient.TTL(key).Result()
	if err == nil && t == -time.Second { // 没有 key 返回 -2s，有 key 但没有 TTL 返回 -1s
		redisClient.Expire(key, ttl)
	}
	return int(result), nil
}

func GetDeleteToken(token string) (bool, error) {
	if token == "" {
		return false, nil
	}
	key := fmt.Sprintf("%v_%v", _PREFIX, token)
	result, err := redisClient.Del(key).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

func PutToken(ip, token string) error {
	if token == "" {
		return nil
	}
	key := fmt.Sprintf("%v_%v", _PREFIX, token)
	return redisClient.Set(key, ip, time.Minute).Err()
}