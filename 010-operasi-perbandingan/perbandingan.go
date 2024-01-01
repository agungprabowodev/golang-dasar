package main

import "fmt"

func main() {
	var name1 = "Agung"
	var name2 = "agung"

	var result bool = name1 != name2
	fmt.Println(result)

	var angka1 = 10
	var angka2 = 12

	var statement bool = angka1 > angka2
	fmt.Println(statement)
}
