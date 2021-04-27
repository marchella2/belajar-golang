package main

import "fmt"

func main() {
	fmt.Println("Belajar Go-Lang")
	fmt.Println("Belajar" + " " + "Go")

	// Cara mengetahui berapa karakter huruf yang digunakan
	fmt.Println(len("Belajar"))

	// Cara untuk mengambil salah satu karakter
	fmt.Println("Belajar"[0])
	fmt.Println("Belajar"[1:4])
	fmt.Println("Belajar"[2:])
	fmt.Println("Belajar"[:4])	
}
