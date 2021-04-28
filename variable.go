package main 

import "fmt"

func main() {
	// Membuat variable diawali dengan var, kemudian nama variable dan terakhir tipe data
	var hello string = "Hello world"

	fmt.Println(hello)

	// mengubah value dari variable yang telah dibuat
	hello = "Halo Golang"

	fmt.Println(hello)

	// cara kedua membuat variable
	var nama string

	fmt.Println(nama) // hasilnya string kosong karena tidak memiliki value

	nama = "Marchella"
	fmt.Println(nama)

	// cara simple membuat variable (tanpa menyebutkan tipe data)
	namaLengkap := "Marchella Putri Sannie"
	fmt.Println(namaLengkap)

	//Membuat variable dengan tipe data integer
	nilai := 10; // tipe data otomatis integer, integer = int32
	fmt.Println(nilai)

	// variable constant = variable yang tidak bisa lagi diubah nilainya
	const perusahaan string = "Programmer zaman now"
	fmt.Println(perusahaan)
}