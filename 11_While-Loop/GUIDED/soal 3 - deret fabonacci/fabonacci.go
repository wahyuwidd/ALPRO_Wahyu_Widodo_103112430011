package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	i := 0
	a, b = 0, 1
	for i < n {
		fmt.Print(a, " ")
		a, b = b, a+b
		i++
	}
}
