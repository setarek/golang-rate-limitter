package redis

import (
	"github.com/go-redis/redis"

	"golang-rate-limitter/config"
)

func GetRedisClient(config *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis_host"),
		Password: config.GetString("redis_password"),
		DB:       config.GetInt("redis_db"),
	})
}
