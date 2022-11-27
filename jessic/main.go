package main

import (
	"fmt"
	"time"
)

func main() {
	var euro float32
	euro = 7.5

	//Ønsker i Euro
	leggins := 9
	leggins2 := 0
	romve_hoodie := 15
	jeans := 17
	jeans2 := 18
	solbriller := 4
	hoodie := 20

	// Ønsker i DKR
	jbl := 749
	cover := 189
	tøjstativ := 706
	airpods := 1150
	kortholder := 101
	øjenvipper := 169
	øjenvipper1 := 89
	jogginbukser := 600
	hoodrich_hoodie := 600
	juicy_hoodie := 680

	resultateuro := leggins + leggins2 + romve_hoodie + jeans + jeans2 + solbriller + hoodie
	resultatdk := jbl + cover + tøjstativ + airpods + kortholder + øjenvipper + øjenvipper1 + jogginbukser + hoodrich_hoodie + juicy_hoodie

	fmt.Println("samlet pris i euro ", resultateuro)

	prisidk := euro * float32(resultateuro)

	total := resultatdk + int(prisidk)

	fmt.Println(total, "Kr")

	time.Sleep(20 * time.Second)
}
