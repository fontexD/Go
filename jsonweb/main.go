package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// The `json:"whatever"` bit is a way to tell the JSON
// encoder and decoder to use those names instead of the
// capitalised names
type person struct {
	Name  string `json:"key"`
	Value string `json:"value"`
}

var tom *person = &person{
	Name:  "status",
<<<<<<< HEAD
	Value: "unHealthy",
=======
	Value: "UNHealthy",
>>>>>>> a6f4352e77f32e2e95c103afd6856374f7a801be
}

func tomHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		// Just send out the JSON version of 'tom'
		j, _ := json.Marshal(tom)
		w.Write(j)
	case "POST":
		// Decode the JSON in the body and overwrite 'tom' with it
		d := json.NewDecoder(r.Body)
		p := &person{}
		err := d.Decode(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		tom = p
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}

func main() {
	http.HandleFunc("/health", tomHandler)

	log.Println("Go!")
	http.ListenAndServe(":8080", nil)
}
