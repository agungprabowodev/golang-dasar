package main

import "fmt"

func main() {
	var i = 10
	i += 10

	fmt.Println(i)

	i += 5

	fmt.Println(i)

	var j = 1
	j++
	fmt.Println(j)
	j++
	fmt.Println(j)
	j--
	fmt.Println(j)
}
