package main

import "fmt"

func main() {
	var bilangan int
	var pernyataan bool
	fmt.Scan(&bilangan)
	if bilangan < 0 && bilangan%2 == 0 {
		pernyataan = true
	}
	fmt.Print(pernyataan)
}
