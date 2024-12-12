package main

import "fmt"

func main() {
	var n int
	var isPrime bool
	isPrime = true
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)
	for j := 2; j*j <= n; j++ {
		if n%j == 0 {
			isPrime = false
		}
	}
	if isPrime {
		fmt.Print("Bilangan Prima")
	} else {
		fmt.Print("Bukan Bilangan Prima")
	}
}
