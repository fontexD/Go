package jsoncrawler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Name  string `json:"key"`
	Value string `json:"status"`
}

func JsonData(urls string) Data {
	var JsonOutput Data
	var err error

	JsonOutput, err = Getdata(urls)
	//error checking
	if err != nil {
		fmt.Println(err)
	}

	return JsonOutput
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
