package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan x dan y (x >= y): ")
	fmt.Scanln(&x, &y)

	hasil := 0
	for x >= y {
		x = x - y
		hasil++
	}

	fmt.Println("Hasil x div y:", hasil)
}
