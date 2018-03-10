package dao

import (
	"fmt"
	"time"
)

func Incr(host string, second int) (int, error) {
	if len(host) > 50 {
		host = host[:50]
	}
	key := fmt.Sprintf("qiniuauth_%v_%v", host, second)

	result, err := redisClient.Incr(key).Result()
	if err != nil {
		return 0, err
	}
	if result == 1 {
		redisClient.Expire(key, time.Duration(second) * time.Second)
	}
	return int(result), nil
}
