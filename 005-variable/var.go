package main

import "fmt"

func main() {
	var name string

	name = "Agung"
	fmt.Println(name)

	name = "Agung Prabowo"
	fmt.Println(name)

	var namalangsung = "Agung"
	fmt.Println(namalangsung)

	namalangsung = "Agung Prabowo"
	fmt.Println(namalangsung)

	namatitikdua := "Agung"
	fmt.Println(namatitikdua)

	namatitikdua = "Agung Prabowo"
	fmt.Println(namatitikdua)

	var (
		namadepan    string
		namabelakang = "Prabowo"
	)

	namadepan = "Agung"
	fmt.Println(namadepan, namabelakang)
}
