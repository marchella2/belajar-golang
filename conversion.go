package main

import "fmt"
import "strconv"

func main() {
	// Konversi tipe data dari int32 ke int64
	var x int32 = 10
	var y int64 = int64(x)

	fmt.Println(y)

	// Konversi tipe data dari int ke float
	var z float64 = float64(y)

	fmt.Println(z)

	// Konversi tipe data dari float ke integer
	var b float64 = 3.3
	var a int32 = int32(b)

	fmt.Println(a)

	// Konversi dari string ke integer
	var nilai string = "100"
	
	nilaiInt, _ := strconv.Atoi(nilai)

	fmt.Println(nilaiInt)

	// Konversi integer ke string
	nilaiString := strconv.Itoa(-100)

	fmt.Println(nilaiString)
}
