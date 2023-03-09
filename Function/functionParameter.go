package main

import (
	"fmt"
)

func namaSaya(firstName string, lastName string) {
	fmt.Println("Halo", firstName, lastName)
}

func main() {
	//pemanggilan model pertama
	namaSaya("Sinung", "Prayetno")

	//model ke2
	firstName := "sinung"
	namaSaya(firstName, "prayetno")
}
