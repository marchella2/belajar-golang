package main

import "fmt"

func main() {
	// pola = nama_variabel := make(map[tipe_data_index]tipe_data_value)
	mahasiswa := make(map[string]string)

	mahasiswa["1001"] = "Marchella" 
	mahasiswa["1002"] = "Citra"
	mahasiswa["1003"] = "Kanaya"  

	fmt.Println(mahasiswa)
	fmt.Println(mahasiswa["1001"])
	fmt.Println(mahasiswa["1002"])
	fmt.Println(mahasiswa["1003"])
	// fmt.Println(mahasiswa["1004"]) //Jika tidak ada isi maka output nya kosong

	for nim, name := range mahasiswa{
		fmt.Println( "NIM", nim, "Bernama", name)
	}
}