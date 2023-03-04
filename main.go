package main

import (
	"fmt"
)

func main() {
	var items []string // Inisialisasi array kosong

	for {
		// Menampilkan pilihan
		fmt.Println("Daftar Produk")
		fmt.Println("1. Senter")
		fmt.Println("2. Radio")
		fmt.Println("3. Antena")
		fmt.Println("4. Obeng")
		fmt.Println("5. Selesai")
		fmt.Print("Pilih nomor item yang akan dimasukkan ke dalam array:")

		// Meminta input dari pengguna
		var choice int
		fmt.Scanln(&choice)

		// Memproses pilihan
		switch choice {
		case 1:
			items = append(items, "Senter")
			fmt.Println("Obeng telah ditambahkan ke dalam array.")
		case 2:
			items = append(items, "Radio")
			fmt.Println("Senter telah ditambahkan ke dalam array.")
		case 3:
			items = append(items, "Antena")
			fmt.Println("Baut telah ditambahkan ke dalam array.")
		case 4:
			items = append(items, "Obeng")
			fmt.Println("Baut telah ditambahkan ke dalam array.")
		case 5:
			fmt.Println("Array items:")
			fmt.Println(items)
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
	}
}
