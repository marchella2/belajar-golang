package main

import "fmt"

func main() {

	// Membuat perulangan tanpa berhenti
	// for {
	// 	fmt.Println("Hello World")
	// }

	// Membuat perulangan dengan suatu kondisi (increment)
	for i := 1; i <= 10; i++ {
		fmt.Println(i, "Hello World")
	}

	fmt.Println("Selesai")

	// Membuat perulangan dengan suatu kondisi (decrement)
	for i := 10; i >= 1; i-- {
		fmt.Println(i, "Hello World")
	}
}
