package main

import "fmt"

func main() {
	var i, alas, tinggi, n int
	var luas float64
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&alas, &tinggi)
		luas = 0.5 * float64(alas*tinggi)
		fmt.Println(luas)
	}
}
