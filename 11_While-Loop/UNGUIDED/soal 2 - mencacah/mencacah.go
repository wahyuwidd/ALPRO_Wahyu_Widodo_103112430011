package main

import "fmt"

func main() {
	var input int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&input)

	for input > 0 {
		digit := input % 10
		fmt.Println(digit)
		input /= 10
	}
}
