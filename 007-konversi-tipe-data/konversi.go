package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println("Nilai32", nilai32)
	fmt.Println("Nilai64", nilai64)
	fmt.Println("Nilai16", nilai16) // Menjadi negatif karena melebihi kapasitas (number overflow) int16 max. 32767

	var nama = "Agung Prabowo"
	var e = nama[0]
	var eString = string(e)

	fmt.Println(nama)
	fmt.Println(e)
	fmt.Println(eString)
}
