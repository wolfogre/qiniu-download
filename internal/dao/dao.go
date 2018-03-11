package dao

import (
	"fmt"
	"time"
)

func Incr(domain string, second int) (int, error) {
	if len(domain) > 50 {
		domain = domain[:50]
	}
	key := fmt.Sprintf("qiniuauth_%v_%v", domain, second)

	result, err := redisClient.Incr(key).Result()
	if err != nil {
		return 0, err
	}
	t, err := redisClient.TTL(key).Result()
	if err == nil && t == -time.Second { // 没有 key 返回 -2s，有 key 但没有 TTL 返回 -1s
		redisClient.Expire(key, time.Duration(second) * time.Second)
	}
	return int(result), nil
}
