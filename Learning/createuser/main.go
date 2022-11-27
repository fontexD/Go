package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type UserData struct {
	Firstname string
	Lastname  string
	Username  string
	Age       int
}
type UserDataJson struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Username  string `json:"Username"`
	Age       int    `json:""Age"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {

	fmt.Println("Please enter The Following Data-forms ")

	var Firstname, Lastname, Username string
	var Age int
	var question1 string
	fmt.Println("Firstname, Lastname, Username, Age")
	inputs, err := fmt.Scanln(&Firstname, &Lastname, &Username, &Age)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("Your input could not be read%v", err)
	}

	userData := UserData{

		Firstname,
		Lastname,
		Username,
		Age,
	}

	fmt.Printf(" you entered a total of %v items into scanln function", inputs)
	fmt.Println("\n----------Loading Data .........")
	time.Sleep(time.Second * 1)
	fmt.Println("\n----------Your Data is as follows .........")
	fmt.Println("\nFullName\tUsername\tAge")
	fmt.Println(userData)

	fmt.Println("Do you want to convert your data to json? ('yes' to do su)")
	fmt.Scan(&question1)

	var mySlice []UserData

	var m1 UserData
	m1.Firstname = userData.Firstname
	m1.Lastname = userData.Lastname
	m1.Username = userData.Username
	m1.Age = userData.Age

	mySlice = append(mySlice, m1)
	if question1 == "yes" {

		encodejson(mySlice)

	}
}

func encodejson(input []UserData) {

	filename := "./data.json"

	f, err := ioutil.ReadFile(filename)
	check(err)

	data2 := []UserDataJson{}
	data := input[0]

	json.Unmarshal(f, &data2)

	dataToAdd := &UserDataJson{

		Firstname: data.Firstname,
		Lastname:  data.Lastname,
		Username:  data.Username,
		Age:       data.Age,
	}

	data2 = append(data2, *dataToAdd)

	dataBytes, err := json.MarshalIndent(data2, "", "     ")
	check(err)

	err = ioutil.WriteFile(filename, dataBytes, 0644)
	check(err)

}
