package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type Redis struct {
	Name   string `json:"Name"`
	Group  string `json:"Group"`
	Env    string `json:"Env"`
	Status string `json:"Status"`
}

type Templatedata struct {
	Name   string `json:"Name"`
	Host   string
	Group  string `json:"Group"`
	Env    string `json:"Env"`
	Status string `json:"Status"`
}

func main() {

	getstate()
}
func getstate() {

	var demo = []Templatedata{
		Templatedata{
			Name:   "Redis",
			Host:   "192.168.10.200",
			Group:  "Services",
			Env:    "Demo",
			Status: "",
		},
		Templatedata{
			Name:   "Redis2",
			Host:   "192.168.10.201",
			Group:  "Services",
			Env:    "Integration",
			Status: "",
		},
	}

	var fs []Redis

	for _, demo2 := range demo {

		a := Redis{Name: demo2.Name, Group: demo2.Group, Env: demo2.Env, Status: RedisHealth(demo2.Host)}
		fs = append(fs, a)
		j, _ := json.Marshal(fs)

		j, _ = json.MarshalIndent(fs, "", "  ")

		fmt.Println(string(j))
	}

}

func RedisHealth(Env string) string {
	env_redis := Env + ":6379"

	// new redis client
	client := redis.NewClient(&redis.Options{

		Addr: env_redis,

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
