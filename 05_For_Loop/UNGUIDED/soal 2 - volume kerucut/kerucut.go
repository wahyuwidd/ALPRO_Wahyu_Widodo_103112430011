package main

import "fmt"

func main() {
	var n, i int
	var r, tinggi, volume float64
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&r, &tinggi)
		volume = 1.0 / 3.0 * 3.14 * r * r * tinggi
		fmt.Println(volume)
	}
}
