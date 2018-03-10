package dao

import "github.com/go-redis/redis"

var (
	redisClient *redis.Client
)

func Init(addr, passport string, db int) error {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: passport,
		DB: db,
	})
	err := client.Ping().Err()
	if err != nil {
		redisClient = client
	}
	return err
}
