package main

import "fmt"

func main() {
	a := 1
	b := 2

	//perbandingan
	result := a == b // sama dengan
	tidak := a != b // tidak sama dengan
	kurang := a < b // kurang dari 
	lebih := a > b // lebih dari
	kurangsama := a <= b // kurang dari atau sama dengan
	lebihsama := a >= b // lebih dari atau sama dengan

	fmt.Println(result)
	fmt.Println(tidak)
	fmt.Println(kurang)
	fmt.Println(lebih)
	fmt.Println(kurangsama)
	fmt.Println(lebihsama)
}