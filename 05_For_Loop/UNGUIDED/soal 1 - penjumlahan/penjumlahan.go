package main

import "fmt"

func main() {
	var n, i int
	var hasil int
	fmt.Print("Input bilangan: ")
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		hasil += i
	}
	fmt.Print("Jumlah dari 1 hingga ", n, " adalah ", hasil)
}
