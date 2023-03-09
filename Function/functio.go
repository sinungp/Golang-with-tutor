package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Halo dunia")
}

func main() {
	for i := 0; i < 10; i++ {
		sayHello()
	}
}
