package checks

import (
	"fmt"

	"github.com/go-redis/redis"
)

func RedisHealth(env string) bool {
	return true
	env_redis := env + ":6379"

	// new redis client
	client := redis.NewClient(&redis.Options{

		Addr: env_redis,

		Password: "",

		DB: 0,
	})

	// test connection

	_, err := client.Ping().Result()
	client.Close()
	fmt.Printf(err.Error())
	return err == nil

	// return pong if server is online
}
