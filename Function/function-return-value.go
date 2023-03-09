package main

import (
	"fmt"
)

//Fungsi memanggil nama
func sayHello(name string) string {
	if name == "" {
		return "Halo bro"
	} else {
		return "Halo, " + name
	}
}

func sayUmur(umur int) string {
	if umur >= 30 {
		return "Kamu tua"
	} else if umur >= 15 {
		return "kamu dewasa"
	} else {
		return "kamu muda"
	}
}

func main() {
	result := sayHello("sinung")
	fmt.Println(result)

	fmt.Println(sayHello("prayetno"))

	hasil := sayHello("")
	fmt.Println(hasil)

	hasilUmur := sayUmur(20)
	fmt.Println(hasilUmur)
}
