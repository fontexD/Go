package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fontexd/go/api/pkg/checks/rabbitmq"
	"github.com/fontexd/go/api/pkg/checks/redis"
	"github.com/fontexd/go/api/pkg/models"
	"github.com/fontexd/go/api/pkg/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var Templatedata models.Templatedata

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")

}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/services/rabbitmq", GetRabbitmq).Methods("GET")
	router.HandleFunc("/services/redis", GetRedis).Methods("GET")
	http.ListenAndServe(":80", router)

}

func GetRabbitmq(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	env := r.URL.Query().Get("env")
	if env == "" {
		env = "Demo"
	}

	fmt.Println("", env)
	id := "Rabbitmq"
	productModel, err := mysql.SqlConn(id, env)
	if err != nil {
		log.Print(err)
	}

	fmt.Println("Getting State")

	products := productModel
	parseData := make([]map[string]interface{}, 0)

	for _, product := range products {
		var singleMap = make(map[string]interface{})
		singleMap["Name"] = product.Name
		singleMap["Type"] = product.Type
		singleMap["Env"] = product.Env
		singleMap["Status"] = rabbitmq.Check(product.Host)

		//		fmt.Println(RedisHealth(demo2.Host))
		parseData = append(parseData, singleMap)

	}

	json.NewEncoder(w).Encode(parseData)

}

func GetRedis(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	env := r.URL.Query().Get("env")

	fmt.Println("", env)
	id := "Redis"
	productModel, err := mysql.SqlConn(id, env)
	if err != nil {
		log.Print(err)
	}

	fmt.Println("Getting State")
	products := productModel
	parseData := make([]map[string]interface{}, 0)

	for _, product := range products {
		var singleMap = make(map[string]interface{})
		singleMap["Name"] = product.Name
		singleMap["Type"] = product.Type
		singleMap["Env"] = product.Env
		singleMap["Status"] = redis.RedisHealth(product.Host)

		//		fmt.Println(RedisHealth(demo2.Host))
		parseData = append(parseData, singleMap)

	}

	json.NewEncoder(w).Encode(parseData)

}
