package main

import "fmt"

func main() {
	type NoKTP string

	var ktpAgung NoKTP = "123456789"

	var contoh string = "111111111"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpAgung)
	fmt.Println(contohKtp)
}
