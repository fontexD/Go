package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Data struct {
	Name  string `json:"key"`
	Value string `json:"value"`
}

func main() {
	Urls := "http://127.0.0.1:8080/health"

	JsonOutput, err := Getdata(Urls)
	if err != nil {
		log.Fatal(err)
		log.Fatal(nil)
	}
	fmt.Println(JsonOutput.Value)

}

func Getdata(Urls string) (Data, error) {
	var JsonOutput Data
	for {
		for range time.Tick(5 * time.Second) {
			rsp, err := http.Get(Urls)
			if err != nil {
				log.Fatal(err)
			}
			defer rsp.Body.Close()
			err = json.NewDecoder(rsp.Body).Decode(&JsonOutput)
			if err != nil {
				log.Fatal(err)
			}

		}

	}

}
