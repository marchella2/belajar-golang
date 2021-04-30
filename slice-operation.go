package main

import "fmt"

func main() {
	// Perbedaan array dengan slice = array harus langsung masukin angka di bagian kurung, sedangkan slice tidak usah
	// slice1 := []string{
	// 	"Marchella",
	// 	"Putri",
	// 	"Sannie",
	// }
	// fmt.Println(slice1)

	// Cara kedua membuat slice
	// nama_slice = make([]tipe_data, value)
	slice1 := make([]string, 3)
	slice1[0] = "Marchella"
	slice1[1] = "Putri"
	slice1[2] =  "Sannie"

	// slice2 := slice1[:]

	// fmt.Println(slice1)

	// slice1[0] = "Budi"
	
	// fmt.Println(slice1)
	// fmt.Println(slice2)

	// cara menambahkan data slice yang sudah ada = menggunakan append , jika menambahkan data slice tidak akan mengubah, tetapi menambah slice baru
	slice2 := append(slice1, "Budi", "Nugraha")
	// operasi append membuat slice dan array baru

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1[0] = "Joko"

	fmt.Println(slice1)
	fmt.Println(slice2)

	// Cara membuat slice baru dengan menggunakan slice yang sudah ada
	// pola = copy (slice_tujuan, slice_sumber)
	slice3 := make([]string, 2)

	copy(slice3, slice1) // hanya akan mencopy data sampai secukupnya

	fmt.Println(slice1)
	fmt.Println(slice3)

	slice1[0] = "Budi"

	fmt.Println(slice1)
	fmt.Println(slice3)
}