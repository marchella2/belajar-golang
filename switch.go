package main

import (
	"fmt"; "runtime"
)


func main() {
	// Penggunaan switch pada integer
	for i := 1; i <= 10; i++ {
		switch i {
		case 1:
			fmt.Println("Satu")
		case 2: 
			fmt.Println("Dua")
		case 3: 
			fmt.Println("Tiga")
		case 4: 
			fmt.Println("Empat")
		case 5: 
			fmt.Println("Lima")
		default: 
			fmt.Println("Gak Tau")
		}
	}

	// Penggunaan switch pada string
	sistemOperasi := runtime.GOOS //untuk mengambil sistem operasi yang digunakan

	switch sistemOperasi{
	case "darwin":
		fmt.Println("Mac")
	case "linux": 
		fmt.Println("Linux")
	default: 
		fmt.Println("Gak Tau")
	}
}