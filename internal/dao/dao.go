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
	if err == nil && t == -1 {
		redisClient.Expire(key, time.Duration(second) * time.Second)
	}
	return int(result), nil
}
