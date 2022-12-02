package main

import (
	"encoding/json"
	"fmt"
	"log"

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
	Host   string `json:"Host"`
	Group  string `json:"Group"`
	Env    string `json:"Env"`
	Status string `json:"Status"`
}

func main() {

	getstate()
}
func getstate() {

	demo := []Templatedata{
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
		Templatedata{
			Name:   "Redis3",
			Host:   "192.168.10.200",
			Group:  "Services",
			Env:    "Test",
			Status: "",
		},
	}

	parseData := make([]map[string]interface{}, 0)
	for _, demo2 := range demo {

		var singleMap = make(map[string]interface{})
		singleMap["Name"] = demo2.Name
		singleMap["Group"] = demo2.Group
		singleMap["Env"] = demo2.Env
		singleMap["Status"] = RedisHealth(demo2.Host)

		parseData = append(parseData, singleMap)
	}
	encodeJson, err := json.Marshal(parseData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(encodeJson))
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
