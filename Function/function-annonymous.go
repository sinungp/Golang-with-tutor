package main

import "fmt"

type Blacklist func(string) bool

func registerName(name string, blacklist Blacklist) {
	if blacklist(name) == true {
		fmt.Println("Kamu di block", name)
	} else {
		fmt.Println("Selamat datang", name)
	}
}

func main() {
	//model pertama
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerName("admin", blacklist)
	registerName("sinung", blacklist)

	//model ke2
	registerName("root", func(name string) bool {
		return name == "root"
	})

	registerName("sinung", func(name string) bool {
		return name == "root"
	})
}
