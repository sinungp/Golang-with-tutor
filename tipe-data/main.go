package main

import (
	"fmt"
)

func main() {
	var nilai32 int32 = 10000
	var nilai16 int16 = int16(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai16)
	fmt.Println(nilai8) //akan terjadi integer overflow maka nilai tidak akan sama pada konversi int 8

	var name = "sinung"
	var s = name[0]
	var eString = string(s)

	fmt.Println(name)
	fmt.Println(eString)
}
