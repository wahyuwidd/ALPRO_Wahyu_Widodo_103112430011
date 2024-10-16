package main

import "fmt"

func main() {
	var a, b, i int
	var pangkat int
	fmt.Scan(&a, &b)
	pangkat = 1
	for i = 1; i <= b; i++ {
		pangkat *= a
	}
	fmt.Print(pangkat)
}
