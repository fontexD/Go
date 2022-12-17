package main

import (
	"fmt"
)

func main() {

	navn := "Peter"

	fmt.Println("Dit navn er: ", navn)

	fmt.Println("Gemt Dette sted i hukommelsen", &navn)
}
