package main

import "fmt"

func sayGoodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	getGoodbye := sayGoodbye
	result := getGoodbye("sinung")
	fmt.Println(result)
}
