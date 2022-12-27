package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fontexd/go/api/pkg/checks/applicationpods"
	"github.com/fontexd/go/api/pkg/checks/kafka"
	"github.com/fontexd/go/api/pkg/checks/rabbitmq"
	"github.com/fontexd/go/api/pkg/checks/redis"
	"github.com/fontexd/go/api/pkg/checks/servicepods"
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
	router.HandleFunc("/services/kafka", GetKafka).Methods("GET")
	router.HandleFunc("/services/pods", GetPods).Methods("GET")
	router.HandleFunc("/applications/pods", GetApplicationPods).Methods("GET")
	http.ListenAndServe(":80", router)

}

func GetApplicationPods(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	env := r.URL.Query().Get("env")

	id := "App"
	productModel, err := mysql.SqlConn(id, env)
	if err != nil {
		log.Print(err)
	}

	products := productModel
	parseData := make([]map[string]interface{}, 0)

	for _, product := range products {
		var singleMap = make(map[string]interface{})
		singleMap["Name"] = product.Name
		singleMap["Type"] = product.Type
		singleMap["Env"] = product.Env
		singleMap["Status"] = applicationpods.AppPodHealthCheck(product.Name, product.Host)

		parseData = append(parseData, singleMap)

	}

	json.NewEncoder(w).Encode(parseData)

}

func GetPods(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	env := r.URL.Query().Get("env")

	id := "Pods"
	productModel, err := mysql.SqlConn(id, env)
	if err != nil {
		log.Print(err)
	}

	products := productModel
	parseData := make([]map[string]interface{}, 0)

	for _, product := range products {
		var singleMap = make(map[string]interface{})
		singleMap["Name"] = product.Name
		singleMap["Type"] = product.Type
		singleMap["Env"] = product.Env
		singleMap["Status"] = servicepods.PodHealthCheck(product.Name, product.Host)

		parseData = append(parseData, singleMap)

	}

	json.NewEncoder(w).Encode(parseData)

}

func GetRabbitmq(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	env := r.URL.Query().Get("env")

	id := "Rabbitmq"
	productModel, err := mysql.SqlConn(id, env)
	if err != nil {
		log.Print(err)
	}

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

func GetKafka(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	env := r.URL.Query().Get("env")
	id := "Kafka"
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
		singleMap["Status"] = kafka.KafkaHealthCheck(product.Host)

		//		fmt.Println(RedisHealth(demo2.Host))
		parseData = append(parseData, singleMap)

	}

	json.NewEncoder(w).Encode(parseData)

}
