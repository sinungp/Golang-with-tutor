package main

import (
	"fmt"
)

func main() {
	var a = 10
	var b = 20
	var penjumlahan = a + b
	var pengurangan = a - b
	var pembagian = a / b
	var perkalian = a * b
	var modulus = a % b
	fmt.Println(penjumlahan)
	fmt.Println(pengurangan)
	fmt.Println(pembagian)
	fmt.Println(perkalian)
	fmt.Println(modulus)

	//Augmented assigment
	a += 10
	fmt.Println(a)
}
