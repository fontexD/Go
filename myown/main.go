package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Name  string `json:"key"`
	Value string `json:"value"`
}

func main() {
	var JsonOutput Data
	Urls := "http://127.0.0.1:8080/health"
	Getdata(Urls)
	fmt.Print(JsonOutput.Value, error)
}

func Getdata(Urls string) (Data, error) {

	var JsonOutput Data

	rsp, err := http.Get(Urls)
	if err != nil {
		log.Fatal(err)
		return JsonOutput, err
	}

	defer rsp.Body.Close()

	err = json.NewDecoder(rsp.Body).Decode(&JsonOutput)
	if err != nil {
		log.Fatal(err)
		return JsonOutput, err
	}
	return JsonOutput, err
}
