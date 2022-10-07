package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Define strucutre of data
type Data struct {
	Name  string `json:"key"`
	Value string `json:"value"`
}

func main() {
	//Define var Jsonoutput to what Datastructure it holds
	var JsonOutput Data
	var err error

	// Create array of urls to send t function GetData
	Urls := []string{"http://127.0.0.1:8080/health", "http://127.0.1.1:8080/health", "http://127.0.2.1:8080/health"}

	// Infinity loop which render every 5 second
loop:
	for x := range Urls {
		JsonOutput, err = Getdata(Urls[x])
		//error checking
		if err != nil {
			log.Fatal(err)
		}
		// print .Value from Jsonoutput Struct variable
		fmt.Println(JsonOutput.Value)
		time.Sleep(5 * time.Second)
		fmt.Print(x)
		fmt.Print(Urls)
		goto loop
	}
}

// Func to pull data from  http.get, it takes Urls as input/call , Returns
func Getdata(Urls string) (Data, error) {

	var JsonOutput Data
	rsp, err := http.Get(Urls)
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Body.Close()

	err = json.NewDecoder(rsp.Body).Decode(&JsonOutput)
	if err != nil {
		log.Fatal(err)
	}
	return JsonOutput, err
}
