package main

import (
	"fmt"
)

func main() {
	var nama, alamat, handphone string
	var umur, nilai_ujian int
	const sekolah = "SMK N 1 Kab Tangerang"

	fmt.Print("Masukan Nama: ")
	fmt.Scan(&nama)
	fmt.Print("Masukan Alamat: ")
	fmt.Scan(&alamat)
	fmt.Print("Masukan Nomor Handphone: ")
	fmt.Scan(&handphone)
	fmt.Print("Masukan Umur: ")
	fmt.Scan(&umur)
	fmt.Print("Masukan Nilai Ujian: ")
	fmt.Scan(&nilai_ujian)

	fmt.Println(sekolah)
	fmt.Println("Nama:", nama)
	fmt.Println("Alamat:", alamat)
	fmt.Println("Nomor Handphone:", handphone)
	fmt.Println("Umur:", umur)
	fmt.Println("Nilai Ujian:", nilai_ujian)
}
