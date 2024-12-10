package main

import (
	"fmt"
)

func isSempurna(n int) bool {
	if n <= 1 {
		return false
	}

	hasil := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			hasil += i
		}
	}

	return hasil == n
}

func main() {
	var bilangan int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scanln(&bilangan)

	if isSempurna(bilangan) {
		fmt.Printf("Ya, %d adalah bilangan sempurna.\n", bilangan)
	} else {
		fmt.Printf("Tidak, %d bukan bilangan sempurna.\n", bilangan)
	}
}
