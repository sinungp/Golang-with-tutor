package main

import (
	"fmt"
)

func main() {
	var name string

	name = "sinung prayetno"
	fmt.Println(name)

	name = "prayetno sinung"
	fmt.Println(name)

	var friendName = "Choirul"
	fmt.Println(friendName)

	var age int
	age = 20
	fmt.Println("Umur saya adalah = ", age)

	var age2 = 20
	fmt.Println("Umur saya adalah = ", age2)

	negara := "indonesia"
	fmt.Println("Negara saya = ", negara)

	negara = "amerika"
	fmt.Println("negara kamu = ", negara)

	var (
		firstName = "sinung"
		lastName  = "prayetno"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
