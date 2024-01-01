package main

import "fmt"

func main() {
	const firstName string = "Agung"
	const lastName = "Prabowo"

	fmt.Println(firstName, lastName)

	// firstName = "Diubah"
	// lastName = "Diubah"

	// fmt.Println(firstName, lastName)

	const (
		namaDepan    = "Agung"
		namaBelakang = "Prabowo"
	)

	fmt.Println(namaDepan, namaBelakang)
}
