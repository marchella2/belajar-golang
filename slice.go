package main

import "fmt"

func main() {
	names := [5]string{
		"Marchella",
		"Putri",
		"Sannie",
		"Halo",
		"Test",
	}

	fmt.Println(names)

	// Slice tidak memiliki data, hanya referensi ke array yang sudah ada (seperti pointer/petunjuk)

	slice1 := names[1:4] // Mengambil data dari data kedua sampai keempat
	fmt.Println(slice1)

	slice2 := names[:3] //Jika yang ingin diambil dimulai dari data pertama maka dikosongkan saja
	fmt.Println(slice2)

	slice3 := names[1:] // Mengambil data dimulai dari satu sampai selesai
	fmt.Println(slice3)

	names[1] = "Nugraha"

	fmt.Println(names)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)	

	slice1[1] = "Moro"
	
	fmt.Println(names)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)	
	//Jika mengubah slice akan merubah value dari array juga
}