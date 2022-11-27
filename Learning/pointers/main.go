package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/user"
	"regexp"
	"time"
)

type userData struct {
	Firstname string
	Lastname  string
	Username  string
	Age       int
	Birthday  time.Time
}

func main() {

	var firstname string
	var age int
	var input string

	textInput := "yes"
	firstname = "Anonymous"
	age = 5

	fmt.Println("Value of firstname is: ", firstname)
	fmt.Println("Current age is: ", age)
	pointerStringChange(&firstname, &age)

	fmt.Println("Loading new values......")

	time.Sleep(3 * time.Second)

	fmt.Println("New value of my string is: ", firstname)
	fmt.Println("New value of my age is: ", age)

	fmt.Println("Would you like enter your own data ? type yes to procced..")

	fmt.Scan(&input)
	inputs(input, firstname, textInput, age)

	if input != textInput {

		user, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		reg := regexp.MustCompile(`^[^\\]*\\`)
		res := reg.ReplaceAllString(user.Username, ``)

		conn, err := net.Dial("ip:icmp", "google.com")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Have a nice day!", res, conn.LocalAddr())
	}
}

func inputs(input string, firstname string, textInput string, age int) {

	fmt.Scan(input)
	var input2 string
	if input == textInput {

		fmt.Println("Enter Value of firstname")
		fmt.Scan(&firstname)
		fmt.Println("Enter Value of age")
		fmt.Scan(&age)

		fmt.Println("Your Value of firstname is: ", firstname, "Your age is: ", age)

		fmt.Println("Do you wantr to store your data to a file ? 'type yes to do so' ")
		fmt.Scan(&input2)
		if input2 == "yes" {
			data1 := []byte(firstname)

			if _, err := os.Stat("/path/to/whatever"); errors.Is(err, os.ErrNotExist) {
				fmt.Println("file exist")
			}
			err := ioutil.WriteFile("data.txt", data1, 0)
			if err != nil {
				log.Fatal(err)
			}

		}
	}
}
func pointerStringChange(a *string, b *int) {

	newValueString := "nowItsChabnged"
	newValueInt := 10
	*a = newValueString
	*b = newValueInt
}
