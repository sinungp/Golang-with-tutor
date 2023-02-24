package main

import "fmt"

var saldo float32 = 1000000

func main() {
	var (
		nama_produk string  = "motor"
		jenis       string  = "kendaraan"
		harga       float32 = 200000
		diskon      float32 = 10
		status      string  = "tersedia"
	)

	var (
		nama_produk2 string  = "baju"
		jenis2       string  = "pakaian"
		harga2       float32 = 50000
		diskon2      float32 = 6
		status2      string  = "tersedia"
	)

	var (
		nama_produk3 string  = "handphone"
		jenis3       string  = "elekronik"
		harga3       float32 = 100000
		diskon3      float32 = 12
		status3      string  = "tersedia"
	)

	//var konfirmasi bool = true
	const ppn = 0.11
	fmt.Println("Nama Produk: ", nama_produk, "Jenis: ", jenis, "Harga: ", harga, "PPN: ", ppn, "Status: ", status)
	fmt.Println("Nama Produk: ", nama_produk2, "Jenis: ", jenis2, "Harga: ", harga2, "PPN: ", ppn, "Status: ", status2)
	fmt.Println("Nama Produk: ", nama_produk3, "Jenis: ", jenis3, "Harga: ", harga3, "PPN: ", ppn, "Status: ", status3)

	fmt.Println("=======================================================================================================")

	hasil_diskon := harga * diskon / 100
	hasil_diskon2 := harga2 * diskon2 / 100
	hasil_diskon3 := harga3 * diskon3 / 100

	hasil_ppn := hasil_diskon * ppn
	hasil_ppn2 := hasil_diskon2 * ppn
	hasil_ppn3 := hasil_diskon3 * ppn

	konfirmasi := "Y"
	fmt.Print("Apakah anda ingin membayar (Y) : ")
	fmt.Scan(&konfirmasi)

	if konfirmasi == "Y" {
		grand_total := (saldo - hasil_diskon - hasil_diskon2 - hasil_diskon3) * ppn
		diskon_total := hasil_diskon + hasil_diskon2 + hasil_diskon3
		harga_total := harga + harga2 + harga3
		ppn_total := hasil_ppn + hasil_ppn2 + hasil_ppn3

		fmt.Println("Total harga barang : ", harga_total)
		fmt.Println("Diskon : ", hasil_diskon, hasil_diskon2, hasil_diskon3)
		fmt.Println("Total Discount : ", diskon_total)
		fmt.Println("Vat : ", ppn*100, "%")
		fmt.Println("Total Vat : ", ppn_total)
		fmt.Println("Grand Total : ", grand_total)
	} else {
		fmt.Println("Input yang anda masukan salah")
	}

	fmt.Println("=======================================================================================================")

	fmt.Print("Mau cek Pendapatan (Y) : ")
	fmt.Scan(&konfirmasi)
	if konfirmasi == "Y" {
		var uang_masuk float32
		grand_total := (saldo - hasil_diskon - hasil_diskon2 - hasil_diskon3) * ppn
		uang_masuk = saldo + grand_total

		if saldo < uang_masuk {
			status = "banyak"
		} else {
			status = "biasa aja"
		}

		fmt.Printf("Uang masuk : %f\n", uang_masuk)
		fmt.Printf("Saldo : %f\n", saldo)
		fmt.Println("Status : ", status)
	}

}
