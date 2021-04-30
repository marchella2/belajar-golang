package main

import "fmt"

func main() {
	var names [5]string

	names[0] = "Marchella"
	names[1] = "Putri"
	names[2] = "Sannie"
	names[3] = "Test"
	names[4] = "Halo"

	// Menampilkan dengan menggunakan perulangan
	for i := 0; i < len(names); i++ { // len(names) berfungsi untuk mengecek berapa panjang dari array
		fmt.Println(names[i])
	}

	// Membuat perulangan kedua (cara simple)
	for _, value := range names { //index = nama variable, nama variable bebas
		// fmt.Println("index", index, "=", value)
		fmt.Println(value) // Jika ingin memanggil value nya
		// Jika variabel index tidak digunakan bisa diganti dengan _
	}

	// fmt.Println(names) // Memanggil semua data array
	// fmt.Println(names[0]) // Memanggil data array yang dipilih

	// Cara membuat array tanpa cara manual
	nama := [5]string{
		"Marchella", 
		"Putri", 
		"Sannie",
		"Halo",
		"Dunia",
	}
	
	fmt.Println(nama);
}
