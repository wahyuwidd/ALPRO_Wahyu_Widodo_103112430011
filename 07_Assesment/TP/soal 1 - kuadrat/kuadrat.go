package main

import "fmt"

func main() {
	var n int
	fmt.Print("Input bilangan: ")
	fmt.Scan(&n)
	// i := 1 karena perulangan mulai dari 1 hingga n
	// perulangan akan berhenti ketika i > n
	for i := 1; i <= n; i++ {
		fmt.Print(i*i, " ") // i*i bentuk perhitungan kuadrat, string kosong agar jadi spasi
	}
}
