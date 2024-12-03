package main

import "fmt"

func main() {
	var n, i int
	fmt.Print("Masukkan bilangan n: ")
	fmt.Scan(&n)
	i = n
	for i > 1 {
		fmt.Print(i, " x ")
		i--
	}
	fmt.Println(1)
}
