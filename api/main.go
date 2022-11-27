package main

import (
	"encoding/json"
	"net/http"

	"github.com/fontexd/go/api/pkg/models"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var env string = "192.168.10.200"
var Templatedata models.Templatedata
var Health = []models.Templatedata{
	{Name: "Redis", Group: "Services", Env: "Production", Status: RedisHealth(env)},
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/services/redis", getstate).Methods("GET")

	http.ListenAndServe(":80", router)
}
func getstate(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Health)

}

func RedisHealth(env string) string {
	env_redis := env + ":6379"

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
