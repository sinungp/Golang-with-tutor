package main

import (
	"fmt"
)

func main() {
	var nilai = 80
	var absen = 100

	var lulusNilai bool = nilai >= 70
	var lulusAbsen bool = absen >= 80

	var lulus bool = lulusAbsen && lulusNilai
	fmt.Println(lulus)
}
