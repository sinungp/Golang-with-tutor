package main

import (
	"fmt"
)

func main() {
	const firstName string = "sinung"
	const lastName = "prayetno"
	const birthday = 1999
	const age int = 20

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(birthday)
	fmt.Println(age)

	const (
		name string = "choirul"
		last        = "ummam"
	)
	fmt.Println(name)
	fmt.Println(last)
}
