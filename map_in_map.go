package main

import "fmt"

func main() {
	// mahasiswa := map[string]string{
	// 	"1001": "Marchella",
	// 	"1002": "Putri",
	// 	"1003": "Sannie",
	// }

	// Pola map dalam map 
	mahasiswa := map[string]map[string]string{
		"1001" : map[string]string{
			"name" : "Marchella",
			"address" : "Indonesia",
			"gender" : "Female",
		},

		"1002" : map[string]string{
			"name" : "Sania",
			"address" : "Indonesia",
			"gender" : "Female",
		},

		"1003" : map[string]string{
			"name" : "Arga",
			"address" : "Indonesia",
			"gender" : "Male",
		},
	}

	// Delete map
	delete(mahasiswa, "1003")

	fmt.Println(mahasiswa)
	// fmt.Println(mahasiswa["1001"]["name"])
}
