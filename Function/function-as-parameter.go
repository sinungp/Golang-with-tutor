package main

import "fmt"

// func sayHelloWithFillter(name string, filter func(string) string) {
// 	fmt.Println("Hello ", filter(name))
// }

// func spamFilter(name string) string {
// 	if name == "Anjing" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }

// func main() {
// 	//pemanggilan model pertama
// 	sayHelloWithFillter("sinung", spamFilter)

// 	//pemanggilan model ke2
// 	filter := spamFilter
// 	sayHelloWithFillter("Anjing", filter)
// }

///////////////////////////////Latihan//////////////////////////////////////////

// func sayUmurWithFilter(umur int, filter func(int) string, hasil func(int) int) {
// 	fmt.Print("kamu termasuk ", filter(umur), "dan umur kamu adalah ", hasil(umur))
// }

// func sayHobi(hobi string) string {
// 	return " Hobi kamu adalah " + hobi
// }

// func filterUmur(umur int) string {
// 	if umur > 30 {
// 		return "tua"
// 	} else if umur > 15 {
// 		return "dewasa"
// 	} else {
// 		return "muda"
// 	}
// }

// func hasilUmur(hasil int) int {
// 	return 23
// }

// func main() {

// 	hobiku := sayHobi("jalan jalan")

// 	sayUmurWithFilter(23, filterUmur, hasilUmur)

// 	fmt.Println(hobiku)
// }
///////////////////////////////Model kedua//////////////////////////////////////////
//Menggunakan type declaration
type Filter func(int) string
type Hasil func(int) int

func sayUmurWithFilter(umur int, filter Filter, hasil Hasil) {
	fmt.Print("kamu termasuk ", filter(umur), "dan umur kamu adalah ", hasil(umur))
}

func sayHobi(hobi string) string {
	return " Hobi kamu adalah " + hobi
}

func filterUmur(umur int) string {
	if umur > 30 {
		return "tua"
	} else if umur > 15 {
		return "dewasa"
	} else {
		return "muda"
	}
}

func hasilUmur(hasil int) int {
	return 23
}

func main() {

	hobiku := sayHobi("jalan jalan")

	sayUmurWithFilter(23, filterUmur, hasilUmur)

	fmt.Println(hobiku)
}
