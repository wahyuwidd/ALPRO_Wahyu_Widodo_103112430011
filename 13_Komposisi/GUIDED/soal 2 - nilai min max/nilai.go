package main

import "fmt"

func main() {
	var a, b, c, min, max int
	fmt.Print("Masukkan 3 bilangan: ")
	fmt.Scan(&a, &b, &c)
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}

	if max < c {
		max = c
	}

	fmt.Println("Terbesar", max)
	fmt.Println("Tekecil", min)
}
