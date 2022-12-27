package redis

import (
	"github.com/go-redis/redis"
)

func RedisHealth(Env string) string {
	env_redis := Env + ":6379"

	// new redis client
	client := redis.NewClient(&redis.Options{

		PoolTimeout: 3,
		Addr:        env_redis,

		Password: "",

		DB: 0,
	})

	// test connection

	_, err := client.Ping().Result()
	if err != nil {
		return "UnHealthy"
	}
	// return pong if server is online
	return "Healthy"
}
