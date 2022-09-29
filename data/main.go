package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Which sites to scrape for health
// Get request

type response struct {
	Name  string `json:"key"`
	Value string `json:"value"`
}

func getData(x string) string {
	var data response
	var data2 string
	resp, err := http.Get(x)
	if err != nil {
		fmt.Print(err)
		return x
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Print(x, data.Value, err)
		return response
	}

	return response
}
func main() {

	x := "http://127.0.0.1:8080/health"
	result := getData(x)
	fmt.Printf(result)
}
