package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://sts.nordpoolgroup.com/connect/token"

	payload := strings.NewReader("grant_type=password&scope=dayahead_api")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic Y2xpZW50X2RheWFoZWFkX2FwaTpjbGllbnRfZGF5YWhlYWRfYXBp")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
