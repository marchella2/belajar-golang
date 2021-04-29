package main

import "fmt"

func main() {
	// continue = berfungsi untuk skip codingan yang berada di bawah continue
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue // nilai genap tidak akan muncul
		}

		fmt.Println(i)

		if i == 50 {
			break // nilai yang dimunculkan hanya sampai 50
		}
	}
}
